import sys

import json
import pytest
import subprocess

from kat.harness import Query
from abstract_tests import AmbassadorTest, HTTP, ServiceType


class IngressStatusTest1(AmbassadorTest):
    status_update = {
        "loadBalancer": {
            "ingress": [{
                "ip": "42.42.42.42"
            }]
        }
    }

    def init(self):
        self.target = HTTP()

    def manifests(self) -> str:
        return super().manifests() + """
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: ambassador
    getambassador.io/ambassador-id: {self.ambassador_id}
  name: {self.name.k8s}
spec:
  rules:
  - http:
      paths:
      - backend:
          serviceName: {self.target.path.k8s}
          servicePort: 80
        path: /{self.name}/
"""

    def queries(self):
        if sys.platform != 'darwin':
            text = json.dumps(self.status_update)

            update_cmd = ['kubestatus', 'Service', '-f', f'metadata.name={self.name.k8s}', '-u', '/dev/fd/0']
            subprocess.run(update_cmd, input=text.encode('utf-8'), timeout=5)

            yield Query(self.url(self.name + "/"))
            yield Query(self.url(f'need-normalization/../{self.name}/'))

    def check(self):
        if sys.platform == 'darwin':
            pytest.xfail('not supported on Darwin')

        for r in self.results:
            if r.backend:
                assert r.backend.name == self.target.path.k8s, (r.backend.name, self.target.path.k8s)
                assert r.backend.request.headers['x-envoy-original-path'][0] == f'/{self.name}/'

        # check for Ingress IP here
        ingress_cmd = ["kubectl", "get", "-o", "json", "ingress", self.path.k8s]
        ingress_run = subprocess.Popen(ingress_cmd, stdout=subprocess.PIPE)
        ingress_out, _ = ingress_run.communicate()
        ingress_json = json.loads(ingress_out)
        assert ingress_json['status'] == self.status_update, f"Expected Ingress status to be {self.status_update}, got {ingress_json['status']} instead"


class IngressStatusTest2(AmbassadorTest):
    status_update = {
        "loadBalancer": {
            "ingress": [{
                "ip": "84.84.84.84"
            }]
        }
    }

    def init(self):
        self.target = HTTP()

    def manifests(self) -> str:
        return super().manifests() + """
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: ambassador
    getambassador.io/ambassador-id: {self.ambassador_id}
  name: {self.name.k8s}
spec:
  rules:
  - http:
      paths:
      - backend:
          serviceName: {self.target.path.k8s}
          servicePort: 80
        path: /{self.name}/
"""

    def queries(self):
        if sys.platform != 'darwin':
            text = json.dumps(self.status_update)

            update_cmd = ['kubestatus', 'Service', '-f', f'metadata.name={self.name.k8s}', '-u', '/dev/fd/0']
            subprocess.run(update_cmd, input=text.encode('utf-8'), timeout=5)

            yield Query(self.url(self.name + "/"))
            yield Query(self.url(f'need-normalization/../{self.name}/'))

    def check(self):
        if sys.platform == 'darwin':
            pytest.xfail('not supported on Darwin')

        for r in self.results:
            if r.backend:
                assert r.backend.name == self.target.path.k8s, (r.backend.name, self.target.path.k8s)
                assert r.backend.request.headers['x-envoy-original-path'][0] == f'/{self.name}/'

        # check for Ingress IP here
        ingress_cmd = ["kubectl", "get", "-o", "json", "ingress", self.path.k8s]
        ingress_run = subprocess.Popen(ingress_cmd, stdout=subprocess.PIPE)
        ingress_out, _ = ingress_run.communicate()
        ingress_json = json.loads(ingress_out)
        assert ingress_json['status'] == self.status_update, f"Expected Ingress status to be {self.status_update}, got {ingress_json['status']} instead"


class IngressStatusTestAcrossNamespaces(AmbassadorTest):
    status_update = {
        "loadBalancer": {
            "ingress": [{
                "ip": "168.168.168.168"
            }]
        }
    }

    def init(self):
        self.target = HTTP(namespace="alt-namespace")

    def manifests(self) -> str:
        return super().manifests() + """
---
apiVersion: v1
kind: Namespace
metadata:
  name: alt-namespace
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: ambassador
    getambassador.io/ambassador-id: {self.ambassador_id}
  name: {self.name.k8s}
  namespace: alt-namespace
spec:
  rules:
  - http:
      paths:
      - backend:
          serviceName: {self.target.path.k8s}
          servicePort: 80
        path: /{self.name}/
"""

    def queries(self):
        if sys.platform != 'darwin':
            text = json.dumps(self.status_update)

            update_cmd = ['kubestatus', 'Service', '-f', f'metadata.name={self.name.k8s}', '-u', '/dev/fd/0']
            subprocess.run(update_cmd, input=text.encode('utf-8'), timeout=5)

            yield Query(self.url(self.name + "/"))
            yield Query(self.url(f'need-normalization/../{self.name}/'))

    def check(self):
        if sys.platform == 'darwin':
            pytest.xfail('not supported on Darwin')

        for r in self.results:
            if r.backend:
                assert r.backend.name == self.target.path.k8s, (r.backend.name, self.target.path.k8s)
                assert r.backend.request.headers['x-envoy-original-path'][0] == f'/{self.name}/'

        # check for Ingress IP here
        ingress_cmd = ["kubectl", "get", "-o", "json", "ingress", self.path.k8s, "-n", "alt-namespace"]
        ingress_run = subprocess.Popen(ingress_cmd, stdout=subprocess.PIPE)
        ingress_out, _ = ingress_run.communicate()
        ingress_json = json.loads(ingress_out)
        assert ingress_json['status'] == self.status_update, f"Expected Ingress status to be {self.status_update}, got {ingress_json['status']} instead"


class IngressStatusTestWithAnnotations(AmbassadorTest):
    status_update = {
        "loadBalancer": {
            "ingress": [{
                "ip": "200.200.200.200"
            }]
        }
    }

    def init(self):
        self.target = HTTP()

    def manifests(self) -> str:
        return super().manifests() + """
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    getambassador.io/config: |
      ---
      apiVersion: ambassador/v1
      kind:  Mapping
      name:  {self.name}-nested
      prefix: /{self.name}-nested/
      service: http://{self.target.path.fqdn}
      ambassador_id: {self.ambassador_id}
    kubernetes.io/ingress.class: ambassador
    getambassador.io/ambassador-id: {self.ambassador_id}
  name: {self.name.k8s}
spec:
  rules:
  - http:
      paths:
      - backend:
          serviceName: {self.target.path.k8s}
          servicePort: 80
        path: /{self.name}/
"""

    def queries(self):
        text = json.dumps(self.status_update)

        update_cmd = ['kubestatus', 'Service', '-f', f'metadata.name={self.name.k8s}', '-u', '/dev/fd/0']
        subprocess.run(update_cmd, input=text.encode('utf-8'), timeout=5)

        yield Query(self.url(self.name + "/"))
        yield Query(self.url(self.name + "-nested/"))
        yield Query(self.url(f'need-normalization/../{self.name}/'))

    def check(self):
        # check for Ingress IP here
        ingress_cmd = ["kubectl", "get", "-o", "json", "ingress", self.path.k8s]
        ingress_run = subprocess.Popen(ingress_cmd, stdout=subprocess.PIPE)
        ingress_out, _ = ingress_run.communicate()
        ingress_json = json.loads(ingress_out)
        assert ingress_json['status'] == self.status_update, f"Expected Ingress status to be {self.status_update}, got {ingress_json['status']} instead"
