# Kubernetes with kind

## How to select context 

```bash
kubectl config use-context NAME
```

## How to create a new cluster
    
```bash
kind create cluster --config=k8s/kind.yaml --name=NAME
```

## How to delete a cluster

```bash
kind delete clusters NAME
```

## How to apply a configuration

```bash
kubectl apply -f k8s/FILE.yaml
```

## How to initialize port 
    
```bash
    kubectl port-forward svc/goserver 8000:8000
```

## test de carga com fortio
    
```bash
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver:8000/healthz"
```