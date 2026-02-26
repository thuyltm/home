The Spark Operator for Apache Spark aims to make specifying and running Spark applications as easy as running other workloads on Kubernetes. It uses Kubernetes custom resources, stores a collection of API objects of a certain kind, for specifying, running and surfacing status of Spark applications

![Architecture](https://www.kubeflow.org/docs/components/spark-operator/overview/architecture-diagram.png)

Specifically, 
1. a user uses the _kubectl_ to create a _SparkApplication_ object. 
2. The __SparkApplication__ controller __receives__ the object through a watcher from the API server, __creates__ a submission carrying the _spark-submit_ arguments, and __sends__ the submission to the _submission runner_. 
3. The __submission runner__ __submits__ the application to run and __creates__ the driver pod of the application. 
4. Upon starting, the driver pod creates the __executor pods__. While the application is running, the __Spark pod monitor__ watches the pods of the application and sends status updates of the pods back to the __controller__, which then updates the status of the application accordingly

#### Experiment
1. Enable Helm is a rerequisite
```sh
microk8s enable helm3
microk8s enable dns
microk8s status
```
2. Add the Helm repository
```sh
microk8s helm repo add --force-update spark-operator https://kubeflow.github.io/spark-operator
microk8s helm repo list
```
3. Install the operator into the spark-operator namespace
```sh
microk8s helm install spark-operator spark-operator/spark-operator --namespace spark-operator --create-namespace --wait
```
Confirm the spark-operator installation 
```sh
% microk8s helm status spark-operator -n spark-operator --show-resources
v1/Service
NAME                         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
spark-operator-webhook-svc   ClusterIP   10.152.183.97   <none>        9443/TCP   44m
v1/Deployment
NAME                        READY   UP-TO-DATE   AVAILABLE   AGE
spark-operator-controller   1/1        1            1        44m
spark-operator-webhook      1/1        1            1        44m
```
4. Create an example application
```sh
microk8s kubectl apply -f my-scala-hello-spark.yaml
```
5. Get the satus of the application
```sh
% microk8s kubectl get sparkapp my-scala-hello-spark
NAME                   SUSPEND   STATUS      ATTEMPTS   START                  FINISH                 AGE
my-scala-hello-spark             COMPLETED   1          2026-02-26T08:55:23Z   2026-02-26T08:55:48Z   31m
# Inspect the detailed pod events and container logs for troubleshooting purposes
% microk8s kubectl describe sparkapp my-scala-hello-spark
Events:
  Type     Reason                        Age                From                          Message
  ----     ------                        ----               ----                          -------
  Warning  SparkApplicationPendingRerun  28s (x3 over 24m)  spark-application-controller  SparkApplication my-scala-hello-spark is pending rerun
  Normal   SparkApplicationSubmitted     25s (x4 over 27m)  spark-application-controller  SparkApplication my-scala-hello-spark was submitted successfully
  Normal   SparkDriverRunning            24s (x4 over 27m)  spark-application-controller  Driver my-scala-hello-spark-driver is running
  Normal   SparkExecutorPending          19s                spark-application-controller  Executor estimatortransformerparamexample-2a3da29c99290e8f-exec-1 is pending
  Normal   SparkExecutorRunning          18s                spark-application-controller  Executor estimatortransformerparamexample-2a3da29c99290e8f-exec-1 is running
  Normal   SparkExecutorCompleted        6s                 spark-application-controller  Executor estimatortransformerparamexample-2a3da29c99290e8f-exec-1 completed
  Normal   SparkDriverCompleted          3s                 spark-application-controller  Driver my-scala-hello-spark-driver completed
% microk8s kubectl logs -f my-scala-hello-spark-driver
```
6. Delete the application
```sh
microk8s kubectl delete sparkapp spark-pi
```