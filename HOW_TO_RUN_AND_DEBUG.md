# Docker

```sh
docker compose -f docker/http/compose.yaml up --build --force-recreate
```

```sh
docker run -it --entrypoint sh http-docker-http-test
docker run --rm gocv/opencv:4.12.0-alpine-ffmpeg-gstreamer go version
docker exec -it opensearch /bin/bash
```

```sh
#remove all stopped containers, unused networks, dangling images, and optionally, volumes
docker system prune
docker system prune --volumes
docker system prune -a
```

```sh
#in the same directory as your compose.yml, this will delete all volumes (as well as containers/network) defined in the configuration file
docker compose down --volumes
```

```sh
# Health Check
docker inspect --format='{{json .State.Health}}' service2
```