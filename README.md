# Kubernetes with kind

## How to create a new cluster
    
```bash
kind create cluster --config=k8s/kind.yaml --name=NAME
```

## How to delete a cluster

```bash
kubectl config delete-cluster kind-NAME
```

## How to apply a configuration

```bash
kubectl apply -f k8s/FILE.yaml
```