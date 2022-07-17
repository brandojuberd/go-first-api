# GO-FIRST-API

HTTP-API build with Go  

## Stacks:
- Gin
- GORM
- Docker
- Kubernetes

## Features:
-[x] Model 'Weapon' with CRUD operations
-[x] Model 'User' with CRUD operations and hooks on password field

## Command:

Run with Kubernetes
1. Set Env
2. Build docker image. Image name should be same on the kubernetes configuration
```console
  docker build . -t brando/go/first-api
```
3. Run the Kubernetes Deployment
```console
  kubectl apply -f infra/app-depl.yaml
```
4. Check the deployment
```console
  kubectl get deployments
```
5. Run the Kubernetes services to access the pods
```console
  kubectl apply -f infra/app-srv.yaml
```
6. Check the services ports
```console
  kubectl get services
  kubectl describe services go-first-api-srv
```
7. Check connection, should return "404 page not found"
```console
  curl localhost:<service-port>
```
