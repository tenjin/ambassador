#!bash

# Coverage checks are totally broken right now. I suspect that it's
# probably the result of all the Ambassador stuff actually happen in
# Docker containers. To restore it, first add
# 
# --cov=ambassador --cov=ambassador_diag --cov-report term-missing
#
# to the pytest line, and, uh, I guess recover and merge all the .coverage 
# files from the containers??

HERE=$(cd $(dirname $0); pwd)
ROOT=$(cd .. ; pwd)

set -e
set -o pipefail
set -x

# We only want to pull images if they are not present locally. This impacts local test runs.
if [[ "$(docker images -q $AMBASSADOR_DOCKER_IMAGE 2> /dev/null)" == "" ]]; then
    if ! docker pull $AMBASSADOR_DOCKER_IMAGE; then
        echo "could not pull $AMBASSADOR_DOCKER_IMAGE" >&2
        exit 1
    fi
fi

if [[ "$USE_KUBERNAUT" != "true" ]]; then
    ( cd "$ROOT"; bash "$HERE/test-warn.sh" )
fi

TEST_ARGS="--tb=short -s"

seq=("")

if [[ -n "${TEST_NAME}" ]]; then
    case "${TEST_NAME}" in
    group0) seq=('Plain' 'not Plain and (A or C)' 'not Plain and not (A or C)') ;;
    group1) seq=('Plain') ;;
    group2) seq=('not Plain and (A or C)') ;;
    group3) seq=('not Plain and not (A or C)') ;;
    *) seq=("$TEST_NAME") ;;
    esac
fi

( cd "$ROOT" ; make cluster-and-teleproxy )

echo "==== [$(date)] ==== STARTING TESTS"

failed=()

for el in "${seq[@]}"; do
    echo "==== [$(date)] ==== running $el"

#    kubectl delete namespaces -l scope=AmbassadorTest
#    kubectl delete all -l scope=AmbassadorTest

    k_args=""

    if [ -n "$el" ]; then
        k_args="-k $el"
    fi

    hr_el="${el:-ALL}"
    outdirbase="kat-log-${hr_el}"
    outdir="/tmp/${outdirbase}"
    tmpdir="/tmp/kat-tmplog-${hr_el}"

    rm -rf "$tmpdir"; mkdir "$tmpdir"

    if ! pytest ${TEST_ARGS} $k_args | tee /tmp/pytest.log; then
        failed+=("$el")

        mv /tmp/pytest.log "$tmpdir"

        kubectl get pods --all-namespaces > "$tmpdir/pods.txt" 2>&1
        kubectl get svc --all-namespaces > "$tmpdir/svc.txt" 2>&1

        if [ -n "${AMBASSADOR_DEV}" ]; then
            docker ps -a > "$tmpdir/docker.txt" 2>&1
        fi

        for pod in $(kubectl get pods -o jsonpath='{range .items[?(@.status.phase != "Running")]}{.metadata.name}:{.status.phase}{"\n"}{end}'); do
            # WTFO.
            podname=$(echo $pod | cut -d: -f1)
            kubectl logs $podname > "$tmpdir/pod-$podname.log" 2>&1
        done
    fi

    if [ -f /tmp/k8s-AmbassadorTest ]; then
        mv /tmp/k8s-AmbassadorTest.yaml "$tmpdir"
    fi

    for file in /tmp/kat-logs-*; do
        if [ "$file" = '/tmp/kat-logs-*' ]; then
            break
        else
            mv $file "$tmpdir"
        fi
    done

    mv /tmp/kat-client* "$tmpdir"
    mv "$tmpdir" "$outdir"

    ( cd /tmp; tar czf "$outdir.tgz" "$outdirbase" )

    now=$(date +"%y-%m-%dT%H:%M:%S")
    branch=${GIT_BRANCH_SANITIZED:-localdev}
    aws_key="kat-${branch}-${now}-${hr_el}-logs.tgz"
    echo "Uploading log tarball as $aws_key"

    aws s3api put-object \
        --bucket datawire-static-files \
        --key kat/${aws_key} \
        --body "${outdir}.tgz"
done

if (( ${#failed[@]} == 0 )); then
    echo "==== [$(date)] ==== FINISHED TESTS (passed)"
    exit 0
else
    echo "==== [$(date)] ==== FINISHED TESTS (failed: ${failed[*]})"
    exit 1
fi

