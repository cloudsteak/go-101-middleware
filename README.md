# go-101-middleware

Middleware Layer of very first Go lang microservice

## Image Creation

### Default

```bash
docker build --tag cloudsteak/go-101-middleware .
```

### Linux image

```bash
docker build --tag cloudsteak/go-101-middleware --platform linux/amd64 .
```

## Run docker image locally

```bash
docker run -d -p 81:4000 --name gomiddleware01 cloudsteak/go-101-middleware:latest
```

### Force delete docker container

```bash
docker rm gomiddleware01 --force
```

## Push image to DocherHub

### Login to DockerHub

```bash
docker login --username=<your DockerHub username>
```

### Tag image

```bash
docker tag $(docker images cloudsteak/go-101-middleware:latest -q) <your DockerHub username>/go-101-middleware:latest
```

### Push image to DockerHub

```bash
docker push <your DockerHub username>/go-101-middleware:latest
```

## Deploy to Kubernetes

### Create namespace

```bash
kubectl create namespace go-101
```

### Deploy web to new namespace

```bash
kubectl apply -f deployment.yaml --namespace go-101
```

### Check deployment status

```bash
kubectl get deployments -n go-101
```

### Check PODs

```bash
kubectl get pods -n go-101
```

### Expose deployemt

```bash
kubectl expose deployment go-101-middleware-deployment --type=NodePort --name=go-101-middleware-svc --target-port=4000 -n go-101
```

### Check nodeport

```bash
kubectl get svc -n go-101
```

## Autoscaling on Kubernetes

### Create

```bash
kubectl autoscale deployment go-101-middleware-deployment --cpu-percent=50 --min=1 --max=10 -n go-101
```

### Check

```bash
kubectl get hpa -n go-101
```

