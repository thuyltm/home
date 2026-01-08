Firstly, build docker image with tag `latest` and push to the Docker Hub
```sh
skaffold build --push
```
Next deploy
```sh
skaffold run
```