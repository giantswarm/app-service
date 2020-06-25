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

- Create cluster and resources.

```bash
kind create cluster
helm install -n app-operator-unique control-plane-catalog/app-operator --version 1.0.3 --dry-run --debug
kubectl apply -f hack/apps.yaml
```

- Check everything is installed.

```bash
helm ls -A
```

```bash
kg get po
```

- Create cluster issuer.

```bash
kubectl apply -f hack/issuer.yaml
```