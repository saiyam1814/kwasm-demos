# WASM IO Workshop

## Setup

- Set up cluster
kind create cluster --image kindest/node:v1.29.0
- Install KWasm
`helm install -n kwasm --create-namespace kwasm-operator kwasm/kwasm-operator --set kwasmOperator.autoProvision=true`
- Check out repository
```
git clone https://github.com/Liquid-Reply/kwasm-demos.git
cd kwasm-demos 
git checkout wasmio-worksop-part1
cd podtato-head
```
- Build podtatohead
```
docker build podtato-head-microservices/crates/podtato-entry -t 0xe282b0/podtato-head-entry --push
docker build podtato-head-microservices/crates/podtato-parts -t 0xe282b0/podtato-head-parts --push
```
- Deploy podtatohead
```sh
helm upgrade \
  --install podtato-head ./delivery/chart \
  --set images.repositoryDirname=0xe282b0 \
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

## Exercise 2: Add Probes
- Startup Probe
- Readiness Probe
- Liveness Probe

## Exercise 3: Add Metrics
- Metrics Endpoint

## Exercise 4: Add Tracing
- Host Functions
- Guest Instrumentation

## Exercise 5: Load Testing
- K6

## Exercise 6: Chaos Testing
- Chaos Mesh
