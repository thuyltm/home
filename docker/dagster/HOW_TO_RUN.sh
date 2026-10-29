#! /bin/sh
eval $(minikube docker-env)
export YOUR_GITHUB_PAT=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
echo $YOUR_GITHUB_PAT | docker login ghcr.io -u thuyltm --password-stdin
docker build -f ./docker/dagster/Dockerfile -t thuyltm2201/my-dagster-quickstart .
docker run -p 4000:4000 thuyltm2201/my-dagster-quickstart:latest


###########################################################################
# Helm
###########################################################################
helm repo add dagster https://dagster-io.github.io/helm
helm repo update

# Install the deployment
helm install dagster dagster/dagster -f values.yaml
kubectl get pods
kubectl port-forward svc/dagster-dagster-webserver 3000:80

kubectl exec -it <your-dagster-pod-name> -- env | grep DAGSTER_PG_PASSWORD
kubectl logs -f <your-dagster-pod-name>

