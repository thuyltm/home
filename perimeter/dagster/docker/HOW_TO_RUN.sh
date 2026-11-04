#! /bin/sh
eval $(minikube docker-env)
export YOUR_GITHUB_PAT=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
echo $YOUR_GITHUB_PAT | docker login ghcr.io -u thuyltm --password-stdin
docker build -f ./perimeter/dagster/docker/Dockerfile -t thuyltm2201/my-dagster-quickstart .
docker run -p 4000:4000 thuyltm2201/my-dagster-quickstart:latest

###########################################################################
# Helm
###########################################################################
helm repo add dagster https://dagster-io.github.io/helm
helm repo update

# Install the deployment
helm install dagster dagster/dagster -f values_local_db.yaml
helm install dagster dagster/dagster -f values_remote_db.yaml
kubectl get pods
kubectl port-forward svc/dagster-dagster-webserver 3000:80

kubectl exec -it <your-dagster-pod-name> -- env | grep DAGSTER_PG_PASSWORD
kubectl logs -f <your-dagster-pod-name>

