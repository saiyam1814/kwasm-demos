# bookinfo demo

## Steps to deploy

1. Set up a kind cluster

```bash
kind create cluster
```

2. Install the kwasm-operator and annotate all nodes for kwasm

```bash
helm repo add kwasm http://kwasm.sh/kwasm-operator/
helm install -n kwasm --create-namespace kwasm-operator kwasm/kwasm-operator
kubectl annotate node --all kwasm.sh/kwasm-node=true
```

3. Create an image pull secret for pulling the wasm images

```bash
kubectl create secret docker-registry ghcr --docker-server=ghcr.io --docker-username="$GITHUB_USERNAME" --docker-password="$GITHUB_PERSONAL_ACCESS_TOKEN" --docker-email="ignore@me.com"
```

4. Deploy the Helm chart. Choose one of the provided value files:

- `values.istio.yaml` deploys the 4 bookinfo components as is from the Istio example
- `values.mixed.yaml` deploys the productinfo dashboard and the ratings service from the Istio image and the details and reviews service as wasm modules
- `values.spin.yaml` deploys all 4 services as wasm modules

```bash
helm upgrade --install bookinfo . -f values.istio.yaml
```

The same command with a different values file can be used to switch between the different combinations.