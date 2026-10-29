The command _eval $(minikube docker-env)_ points your local terminal's Docker CLI to the Docker daemon running inside the Minikube cluster

Instead of building images on your host machine and needing to push them to an external registry (like Docker Hub), this command configures your current shell to interact directly with Minikube's internal container engine.