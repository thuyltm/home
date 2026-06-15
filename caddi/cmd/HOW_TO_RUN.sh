1> Place a call through the Gateway IP address named caddi-cmd-gateway 
```sh
curl -v -HHost:caddi.cmd.com --resolve "caddi.cmd.com:443:10.96.221.121" --cacert server.pem "https://caddi.cmd.com:443/caddi-cmd/ping"
```

2> Route the call through the istio-ingress/istio-ingress LoadBalancer
```sh
k get services -n istio-ingress
#NAME            TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                                      AGE
#istio-ingress   LoadBalancer   10.100.30.107   10.100.30.107   15021:30918/TCP,80:31685/TCP,443:31190/TCP   31m
curl -v  -H "Host: caddi.cmd.com" http://10.96.40.235:80/caddi-cmd/ping
curl -v  -H "Host: caddi.cmd.com" --tlsv1.3 --cacert server.pem https://10.96.40.235:443/caddi-cmd/ping
curl -v -HHost:caddi.cmd.com --resolve "caddi.cmd.com:443:10.96.221.121" --cacert server.pem "https://caddi.cmd.com:443/caddi-cmd/ping"
```

curl -v -HHost:caddi.cmd.com --resolve "caddi.cmd.com:443:10.100.30.107" --cacert server.pem "https://caddi.cmd.com:443/caddi-cmd/ping"

3> Test the raw TCP connection:bash
```sh
openssl s_client -connect 10.96.40.235:443 -tls1_3
```