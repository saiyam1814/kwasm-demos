# WASM IO Workshop

## Setup

- [ ] Set up a cluster
- [ ] Configure a container registry
- [ ] Install KWasm
```sh
helm install -n kwasm --create-namespace kwasm-operator kwasm/kwasm-operator --set kwasmOperator.autoProvision=true
```
- Check out this repository
```sh
git clone https://github.com/Liquid-Reply/kwasm-demos.git
```
- Deploy podtatohead
```sh
helm upgrade \
  --install podtato-head ./delivery/chart
```
- Build podtatohead Wasm
```sh
docker build podtato-head-microservices/crates/podtato-entry -t <REGISTRY> /podtato-head-entry --push
docker build podtato-head-microservices/crates/podtato-parts -t <REGISTRY> /podtato-head-parts --push
```
- Deploy podtatohead Wasm
```sh
helm upgrade \
  --install podtato-head ./delivery/chart \
  --set images.repositoryDirname=<REGISTRY> \
  --set images.pullPolicy=Always \
  --set entry.repositoryBasename=podtato-head-entry \
  --set entry.runtimeClassName=wasmedge \
  --set hat.repositoryBasename=podtato-head-parts \
  --set hat.runtimeClassName=wasmedge \
  --set leftLeg.repositoryBasename=podtato-head-parts \
  --set leftLeg.runtimeClassName=wasmedge \
  --set leftArm.repositoryBasename=podtato-head-parts \
  --set leftArm.runtimeClassName=wasmedge \
  --set rightLeg.repositoryBasename=podtato-head-parts \
  --set rightLeg.runtimeClassName=wasmedge \
  --set rightArm.repositoryBasename=podtato-head-parts \
  --set rightArm.runtimeClassName=wasmedge
```

## Exercise 1: Add Logging

## Exercise 2: Health Checks
- [ ] Startup Probe
- [ ] Readiness Probe
- [ ] Liveness Probe

## Exercise 3: Add Metrics
- Metrics Endpoint

## Exercise 4: Add Tracing
- https://github.com/open-telemetry/opentelemetry-rust/issues/1478
- 

## Exercise 5: Load Testing
- K6
```sh
helm repo add grafana https://grafana.github.io/helm-charts
helm install k6-operator grafana/k6-operator
```

## Exercise 6: Chaos Testing
- Chaos Mesh
```sh
helm repo add chaos-mesh https://charts.chaos-mesh.org
helm install chaos-mesh chaos-mesh/chaos-mesh -n=chaos-mesh --set chaosDaemon.runtime=containerd --set chaosDaemon.socketPath=/run/k3s/containerd/containerd.sock --version 2.6.3 --create-namespace
```