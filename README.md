[Specifying external dependencies](https://github.com/bazel-contrib/rules_go/blob/master/docs/go/core/bzlmod.md#external-dependencies)<br/>


it is recommended to manage Go dependencies via go.mod. The go_deps extension parses this file directly, so external tooling such as gazelle update-repos is no longer needed.<br/>

(re-)generate BUILD files. <br/>

```sh 
bazel run //:gazelle 
```

An initial go.mod file can be created via

```sh
bazel run @rules_go//go mod init home
```

A dependency can be added via

```sh
bazel run @rules_go//go get github.com/labstack/echo/v4
```

Run

```sh
bazel run @rules_go//go run bamboo/server/main.go
```

```sh
docker compose -f docker/http/compose.yaml up --build --force-recreate
```

```sh
docker run -it --entrypoint sh http-docker-http-test
docker run --rm gocv/opencv:4.12.0-alpine-ffmpeg-gstreamer go version
```

```sh
#remove all stopped containers, unused networks, dangling images, and optionally, volumes
docker system prune
docker system prune --volumes
```