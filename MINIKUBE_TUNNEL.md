To rectify this problem
```sh
telnet 10.104.233.238 80
Trying 10.104.233.238...
telnet: Unable to connect to remote host: No route to host
```
**minikube tunnel** is a commandline utility that creates a network route on your host machine to allow direct access to K9s services deployed with the LoadBalancer type