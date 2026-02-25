### Submit task to Kubernete

#### Build a fat Jar in Scala
TO build a fat Jar in Scala, the standard and recommended method is to use the __sbt-assembly plugin__. This plogin combines your compiled application code, all its dependencies and the Scala runtime libraries into a single, self-container JAR file that can be run anywhere
```sh
sbt -J-Xmx2G -J-XX:+UseG1GC clean assembly
```

Prefixing the master string with __--master k8s://https://<k8s-apiserver-host>:<k8s-apiserver-port>__ will cause the Spark application to launch on the Kubernetes cluster, with invoking the API server at _<k8s-apiserver-host>_
```sh
% microk8s kubectl config view | grep server
server: https://127.0.0.1:16443
% microk8s kubectl cluster-info
Kubernetes control plane is running at https://127.0.0.1:16443
CoreDNS is running at https://127.0.0.1:16443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
```

#### Build a Docker image that can be deployed into the containers within pods
Using the built-in Microk8s Registry (Recommended for local deployment)
1. Enable the registry addon
```sh
microk8s enable registry
```
2. Tag your Docker image

You need to retag your locally built Docker image to point to the local registry endpoint (localhost:32000)
```sh
% docker build -t localhost:32000/hello-spark:latest .
% docker push localhost:32000/hello-spark:latest
# check whether hello-spark-assembly-1.0.jar existed in the folder /opt/spark/jars
% docker run -it localhost:32000/hello-spark:latest ls /opt/spark/jars | grep "spark"
```

#### Submit Applications to Kubernetes
```sh
spark-submit \
  --class "SimpleApp" \
  --master k8s://https://127.0.0.1:16443 \
  --deploy-mode cluster \
  --name hello-spark-app-name \
  --conf spark.kubernetes.container.image=localhost:32000/hello-spark:latest --conf spark.kubernetes.authenticate.driver.serviceAccountName=spark --conf spark.kubernetes.file.upload.path=local:///opt/spark/tmp  --conf spark.kubernetes.authenticate.submission.caCertFile=microk8s-ca.crt --conf spark.kubernetes.authenticate.submission.oauthToken=$(microk8s kubectl create token spark -n default) \
  local:///opt/spark/jars/hello-spark-assembly-1.0.jar
```

  
__RBAC__

The Spark driver pod uses a Kubernetes service account to access the Kubernetes API server to create and watch executor pods. Specifically, at minimum, the service account must be granted a Role or ClusterRole that allows driver pods to create pods and services.

To create a custom service account, a user can use the below command:
```sh
microk8s kubectl create serviceaccount spark
```

To grant a service account a Role or ClusterRole, a RoleBinding or ClusterRoleBinding is needed. 
```sh
microk8s kubectl create clusterrolebinding spark-role --clusterrole=edit --serviceaccount=default:spark --namespace=default
```
Note that a Role can only be used to grant access to resources (like pods) within a single namespace, whereas a ClusterRole can be used to grant access to cluster-scoped resources (like nodes) as well as namespaced resources accross all namespaces

Verify the spark service account has a permission to create pods in microk8s
```sh
% microk8s kubectl auth can-i create pods --as=system:serviceaccount:default:spark
yes
```

#### Microk8s certificates
Microk8s automatically generates self-signed certificates for _localhost_. The certificate and key files are stored in the _/var/snap/microk8s/current/certs/_ directoy on your system

You can use the certificates with tools like curl to interact with the Kubernetes API server securely
- Certificate Authority (CA) certificate: ca.crt
- Server certificate: server.crt
- Server key: server.key
```sh
% curl -L --cert ${SNAP_DATA}/certs/server.crt --key ${SNAP_DATA}/certs/server.key --cacert ${SNAP_DATA}/certs/ca.crt https://127.0.0.1:16443/readyzz
ok (base)
```

microk8s kubectl cp default/api-test-pod:/run/secrets/kubernetes.io/serviceaccount/..2026_02_25_05_22_39.3440028343/ca.crt ~/Documents/Learn/home/perimeter/spark/hello-spark/ca_2.crt