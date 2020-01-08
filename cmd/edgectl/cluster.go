package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"syscall"
	"time"

	"github.com/pkg/errors"

	"github.com/datawire/ambassador/pkg/supervisor"
)

// Connect the daemon to a cluster
func (d *Daemon) Connect(p *supervisor.Process, out *Emitter, rai *RunAsInfo, kargs []string) error {
	// Sanity checks
	if d.cluster != nil {
		out.Println("Already connected")
		return nil
	}
	if d.bridge != nil {
		out.Println("Not ready: Trying to disconnect")
		return nil
	}
	if !d.network.IsOkay() {
		out.Println("Not ready: Establishing network overrides")
		return nil
	}

	out.Println("Connecting...")
	cluster, err := TrackKCluster(p, rai, kargs)
	if err != nil {
		out.Println(err.Error())
		out.SendExit(1)
		return nil
	}
	d.cluster = cluster

	bridge, err := CheckedRetryingCommand(
		p,
		"bridge",
		[]string{edgectl, "teleproxy", "bridge"},
		rai,
		checkBridge,
		15*time.Second,
	)
	if err != nil {
		out.Println(err.Error())
		out.SendExit(1)
		d.cluster.Close()
		d.cluster = nil
		return nil
	}
	d.bridge = bridge
	d.cluster.SetBridgeCheck(d.bridge.IsOkay)

	out.Printf(
		"Connected to context %s (%s)\n", d.cluster.Context(), d.cluster.Server(),
	)

	tmgr, err := NewTrafficManager(p, d.cluster)
	if err != nil {
		out.Println()
		out.Println("Unable to connect to the traffic manager in your cluster.")
		out.Println("The intercept feature will not be available.")
		out.Println("Error was:", err)
		// out.Println("Use <some command> to set up the traffic manager.") // FIXME
	} else {
		d.trafficMgr = tmgr
	}
	return nil
}

// Disconnect from the connected cluster
func (d *Daemon) Disconnect(p *supervisor.Process, out *Emitter) error {
	// Sanity checks
	if d.cluster == nil {
		out.Println("Not connected")
		return nil
	}

	_ = d.ClearIntercepts(p)
	if d.bridge != nil {
		d.cluster.SetBridgeCheck(nil) // Stop depending on this bridge
		_ = d.bridge.Close()
		d.bridge = nil
	}
	if d.trafficMgr != nil {
		_ = d.trafficMgr.Close()
		d.trafficMgr = nil
	}
	err := d.cluster.Close()
	d.cluster = nil

	out.Println("Disconnected")
	return err
}

// checkBridge checks the status of teleproxy bridge by doing the equivalent of
// curl -k https://kubernetes/api/. It's okay to create a new client each time
// because we don't want to reuse connections.
func checkBridge(p *supervisor.Process) error {
	// A zero-value transport is (probably) okay because we set a tight overall
	// timeout on the client
	tr := &http.Transport{
		// #nosec G402
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Timeout: 5 * time.Second, Transport: tr}
	res, err := client.Get("https://kubernetes.default/api/")
	if err != nil {
		return errors.Wrap(err, "get")
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return errors.Wrap(err, "read body")
	}
	return nil
}

// GetFreePort asks the kernel for a free open port that is ready to use.
// Similar to telepresence.utilities.find_free_port()
func GetFreePort() (int, error) {
	lc := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var operr error
			fn := func(fd uintptr) {
				operr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
			}
			if err := c.Control(fn); err != nil {
				return err
			}
			return operr
		},
	}
	l, err := lc.Listen(context.Background(), "tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// TrafficManager is a handle to access the Traffic Manager in a
// cluster.
type TrafficManager struct {
	crc            Resource
	apiPort        int
	sshPort        int
	client         *http.Client
	interceptables []string
	totalClusCepts int
}

// NewTrafficManager returns a TrafficManager resource for the given
// cluster if it has a Traffic Manager service.
func NewTrafficManager(p *supervisor.Process, cluster *KCluster) (*TrafficManager, error) {
	cmd := cluster.GetKubectlCmd(p, "get", "svc/telepresence-proxy", "deploy/telepresence-proxy")
	err := cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, "kubectl get svc/deploy telepresency-proxy")
	}

	apiPort, err := GetFreePort()
	if err != nil {
		return nil, errors.Wrap(err, "get free port for API")
	}
	sshPort, err := GetFreePort()
	if err != nil {
		return nil, errors.Wrap(err, "get free port for ssh")
	}
	kpfArgStr := fmt.Sprintf("port-forward svc/telepresence-proxy %d:8022 %d:8081", sshPort, apiPort)
	kpfArgs := cluster.GetKubectlArgs(strings.Fields(kpfArgStr)...)
	tm := &TrafficManager{apiPort: apiPort, sshPort: sshPort}

	pf, err := CheckedRetryingCommand(p, "traffic-kpf", kpfArgs, cluster.RAI(), tm.check, 15*time.Second)
	if err != nil {
		return nil, err
	}
	tm.crc = pf
	tm.client = &http.Client{Timeout: 3 * time.Second}
	return tm, nil
}

func (tm *TrafficManager) check(p *supervisor.Process) error {
	body, _, err := tm.request("GET", "state", []byte{})
	if err != nil {
		return err
	}
	var state map[string]interface{}
	if err := json.Unmarshal([]byte(body), &state); err != nil {
		p.Logf("check: bad JSON from tm: %v", err)
		p.Logf("check: JSON data is: %q", body)
		return err
	}
	deployments, ok := state["Deployments"].(map[string]interface{})
	if !ok {
		p.Log("check: failed to get deployment info")
		p.Logf("check: JSON data is: %q", body)
	}
	tm.interceptables = make([]string, len(deployments))
	tm.totalClusCepts = 0
	idx := 0
	for deployment := range deployments {
		tm.interceptables[idx] = deployment
		idx++
		info, ok := deployments[deployment].(map[string]interface{})
		if !ok {
			continue
		}
		cepts, ok := info["Intercepts"].([]interface{})
		if ok {
			tm.totalClusCepts += len(cepts)
		}
	}
	return nil
}

func (tm *TrafficManager) request(method, path string, data []byte) (result string, code int, err error) {
	url := fmt.Sprintf("http://127.0.0.1:%d/%s", tm.apiPort, path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	resp, err := tm.client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "get")
		return
	}
	defer resp.Body.Close()
	code = resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "read body")
		return
	}
	result = string(body)
	return
}

// Name implements Resource
func (tm *TrafficManager) Name() string {
	return "trafficMgr"
}

// IsOkay implements Resource
func (tm *TrafficManager) IsOkay() bool {
	return tm.crc.IsOkay()
}

// Close implements Resource
func (tm *TrafficManager) Close() error {
	return tm.crc.Close()
}
