cat ~/.docker/config.json | base64 -w 0

kubectl create secret generic docker-secret --from-file=.dockerconfigjson=/home/thuy/.docker/config.json --type=kubernetes.io/dockerconfigjson