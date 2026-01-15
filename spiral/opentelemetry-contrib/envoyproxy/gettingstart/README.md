[Guide](https://www.envoyproxy.io/docs/envoy/latest/start/docker)

By default, the Envoy OCI image will start as the root user but will switch to the _envoy_ user created at build time, in the Docker _ENTRYPONT_

The _envoy_ user also needs to have permission to access any required configuration files mounted into the container. Changing the _uid_ and/or _gid_ of the _envoy_ user inside the container. The default _uid_ and _gid_ for the _envoy_ user are _101_

The _uid_ and _gid_ of this user can be set at runtime using the _ENVOY_UID_ and _ENVOY_GID_ environment variables

```sh
mkdir logs
chown 777 logs
docker run -d --name envoy -v $(pwd)/logs:/var/log -e ENVOY_UID=777 -e ENVOY_GID=777 envoyproxy/envoy:dev-6c261a8848fc4ca3c3906c848db99dd3e3dbc33f --log-path logs/custom.log
```

To run the process inside the container as the _root_ user you can set _ENVOY_UID_ to _0_, but doing so has the potential to weaken the security of your running container.

One method of doing this without changing any file permissions is to start the container with the host user's _uid_
```sh
docker run -d --name envoy -v $(pwd)/envoy.yaml:/etc/envoy/envoy.yaml -e ENVOY_UID=$(id -u) envoyproxy/envoy:dev-6c261a8848fc4ca3c3906c848db99dd3e3dbc33f
```


