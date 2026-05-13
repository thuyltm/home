To resolve the error while connecting to the Docker API

"
Not healthy: "docker version --format {{.Server.Os}}-{{.Server.Version}}:{{.Server.Platform.Name}}" exit status 1: permission denied while trying to connect to the docker API at unix:///var/run/docker.sock
",

you must add your current user to the docker group.  Only the root user or members of this group have the permissions requored to access the Docker Unix socket at /var/run/docker.sock

Follow these steps to grant your user the necessary permissions permanently
1. Create the docker group
```sh
sudo groupadd docker
```
2. Add your user to the group
```sh
sudo usermod -aG docker $USER
```
3. Apply the changes to take effect, you should log out and log back in