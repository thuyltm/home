**Generate the certificates and keys used by the Gateway to terminate client TLS connections**
1. Create a root certificate and private key to sign certificates

A root certificate is a self-signed certificate that is the top-level "trust anchor" in a public key infrastructure (PKI) hierarchy
```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=example Inc./CN=example.com' -keyout example.com.key -out example.com.crt
#Output
example.com.crt 
example.com.key
```
2. Create a certificate and a private key for www.example.com

```sh
openssl req -out www.example.com.csr -newkey rsa:2048 -nodes -keyout www.example.com.key -subj "/CN=www.example.com/O=example organization"
# Output
www.example.com.csr # a certificate signing request
www.example.com.key # private key
```
Create a self-signed certificate by using the CSR and private key

A certicate (or end-entity certificate) is a digital document issued to a user, server, or device, which is signed by root certificates
```sh
openssl x509 -req -days 365 -CA example.com.crt -CAkey example.com.key -set_serial 0 -in www.example.com.csr -out www.example.com.crt
# Output
www.example.com.crt
```
3. Store the cert/key in a Secret Kubernete
```sh
kubectl create secret tls example-cert --key=www.example.com.key --cert=www.example.com.crt
```
4. Update the Gateway to include an HTTPS listener that listens on port 443 and references the _example-cert_ Secret Kubernete
```sh
kubectl patch gateway eg --type=json --patch '
  - op: add
    path: /spec/listeners/-
    value:
      name: https
      protocol: HTTPS
      port: 443
      tls:
        mode: Terminate
        certificateRefs:
        - kind: Secret
          group: ""
          name: example-cert
  '
```
Output
```sh
spec:
  gatewayClassName: eg
  infrastructure:
    parametersRef:
      group: gateway.envoyproxy.io
      kind: EnvoyProxy
      name: graceful-shutdown-config
  listeners:
  - allowedRoutes:
      namespaces:
        from: Same
    name: http
    port: 80
    protocol: HTTP
  - allowedRoutes:
      namespaces:
        from: Same
    name: https
    port: 443
    protocol: HTTPS
    tls:
      certificateRefs:
      - group: ""
        kind: Secret
        name: example-cert
      mode: Terminate
```

for i in {1..4}; do curl -I --resolve "www.example.com:80:$( echo $GATEWAY_HOST )" --header "x-user-id: one" www.example.com/serviceb; sleep 1; done

curl -v -HHost:www.example.com --resolve "www.example.com:8443:127.0.0.1" \
--cacert example.com.crt https://www.example.com:8443/servicea

5. Testing using port-forward method of service envoy-default-eg-....
    1. Get service envoy-default-eg-....
```sh
kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-namespace=default,gateway.envoyproxy.io/owning-gateway-name=eg -o jsonpath='{.items[0].metadata.name}'
```
  2. Port-forward to the Envoy service    
```sh
kubectl -n envoy-gateway-system port-forward service/${ENVOY_SERVICE} 8443:443 &
```
  3. Test
```sh
curl -v -HHost:www.example.com --resolve "www.example.com:8443:127.0.0.1" --cacert elliptic/envoy/tls/example.com.crt https://www.example.com:8443/serviceb
```
Output
```sh
TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
```

6. Test through Gateway
```sh
export GATEWAY_HOST=$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}')
curl -v --resolve "www.example.com:443:${GATEWAY_HOST}" --cacert elliptic/envoy/tls/example.com.crt https://www.example.com:443/serviceb
```
