1. Netcat check

Run one of these belowed commands from another node to see if the port is open and reachable
```sh
nc -zv <ip> <port>
telnet <ip> <port>
```

2. Socket Status

Run one of these belowed commands on the local Vault host to verify the process is actively listening on the configured bind address
```sh
ss -tlnp | grep 8201
netstat -an | grep 8201
```