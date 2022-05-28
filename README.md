# Introductionto k8s-resource-finder

k8s-resource-finder is a simple GO app to look for unused resources, namely, secrets, configmaps, service accounts and persistentvolumeclaims.

## Usage

Make sure you have select the correct Kubernetes context:

```bash
go build -o k8s-resource-finder && ./k8s-resource-finder --namespaces monitoring --namespaces default
```

k8s-resource-finder will not change or delete any resource. It is just viewing which pods use what resources.