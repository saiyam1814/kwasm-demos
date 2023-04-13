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

3. Deploy the Helm chart. Choose one of the provided value files:

- `chart/values.istio.yaml` deploys the 4 bookinfo components as is from the Istio example
- `chart/values.mixed.yaml` deploys the productinfo dashboard and the ratings service from the Istio image and the details and reviews service as wasm modules
- `chart/values.spin.yaml` deploys all 4 services as wasm modules

```bash
helm upgrade --install bookinfo ./chart -f chart/values.istio.yaml
```

The same command with a different values file can be used to switch between the different combinations.

4. (Create an image pull secret for pulling the wasm images)

Should not be necessary any more, but in case we have set the package visibility to private again, you can try creating an image pull secret:

```bash
kubectl create secret docker-registry ghcr --docker-server=ghcr.io --docker-username="$GITHUB_USERNAME" --docker-password="$GITHUB_PERSONAL_ACCESS_TOKEN" --docker-email="ignore@me.com"
```
