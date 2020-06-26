# Local Dev

- Docs for installing app-service in KIND using app-operator.
- Please make sure you have Helm 3 installed with the control-plane-catalog.
- See [docs](https://intranet.giantswarm.io/docs/dev-and-releng/helm/).

## Resources

See [apps.yaml](hack/apps.yaml)

- control-plane-catalog appcatalog CR
- control-plane-catalog-configmap CR
- chart-operator-unique app CR (bootstraps chart-operator)
- g8s-cert-manager-unique app CR
- selfsigned-giantswarm clusterissuer CR

## Install

- Create cluster.

```bash
kind create cluster --config hack/kind-config.yaml
```

- Install resources.

```bash
kubectl create namespace giantswarm
kubectl create priorityclass giantswarm-critical
helm install -n giantswarm app-operator-unique control-plane-catalog/app-operator
helm install -n giantswarm chart-operator-unique control-plane-catalog/chart-operator
```

- Install upstream cert-manager (TODO: Switch to g8s-cert-manager).

```bash
kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.14.1/cert-manager.yaml
```

- Create cluster issuer once cert-manager webhook is running.

```bash
kubectl apply -f hack/issuer.yaml
```
- Build docker image.

```bash
CGO_ENABLED=0 GOOS=linux go build .
docker build . -t quay.io/giantswarm/app-service:local-dev
kind load docker-image quay.io/giantswarm/app-service:local-dev
```

- Install Helm chart. (TODO: Script manual changes to remove architect templating.) 

```bash
helm install -n giantswarm app-service-unique ./helm/app-service -f ./helm/app-service/ci/default-values.yaml
```

- Check logs and create some app CRs! 

```bash
kubectl -n giantswarm logs -f deploy/app-service-unique | luigi
```

## Clean Up

- Either `kind delete cluster` or follow these steps.

```bash
kubectl delete -f hack/issuer.yaml
helm -n giantswarm del app-operator-unique app-service-unique chart-operator-unique
kubectl delete priorityclass giantswarm-critical
kubectl delete namespace giantswarm
```
