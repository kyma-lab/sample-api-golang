#Sample Kyma container with go lang

Download kubeconfig.yml and export:
```shell script
export KUBECONFIG=~/Downloads/kubeconfig.yml   
```

```shell script
go run apiServer.go  
```

```shell script
docker build -t kymalab/sample-book-api -f ./docker/Dockerfile . 
```
```shell script
docker run -p 80:8000 kymalab/sample-book-api
```

```shell script
docker push kymalab/sample-book-api  
```

```shell script
kubectl create namespace sample-api
```
if you want to use a different namespace, you need to change the 3 yaml files under k8s

```shell script
kubectl apply -f ./k8s/deployment.yaml 
kubectl apply -f ./k8s/service.yaml 
kubectl apply -f ./k8s/apirule.yaml 
```
#Insperation & Links
- https://blogs.sap.com/2020/11/23/kyma-for-dymmies-first-simple-app-in-sap-cloud-platform-kyma-runtime/
- https://sap-samples.github.io/cloud-cap-risk-management/Kyma/


