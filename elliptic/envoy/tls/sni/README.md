1. Create a root certificate and private key to sign certificates

A root certificate is a self-signed certificate that is the top-level "trust anchor" in a public key infrastructure (PKI) hierarchy
```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=sample Inc./CN=sample.com' -keyout sample.com.key -out sample.com.crt
#Output
sample.com.crt 
sample.com.key
```
2. Create a certificate and a private key for www.sample.com

```sh
openssl req -out www.sample.com.csr -newkey rsa:2048 -nodes -keyout www.sample.com.key -subj "/CN=www.sample.com/O=sample organization"
# Output
www.sample.com.csr # a certificate signing request
www.sample.com.key # private key
```
Create a self-signed certificate by using the CSR and private key

3. A certicate (or end-entity certificate) is a digital document issued to a user, server, or device, which is signed by root certificates

Triple: sample.com.crt (CA) + sample.com.key (CAkey) + www.sample.com.csr (A certificate signing request) = A end-entity certificate
```sh
openssl x509 -req -days 365 -CA sample.com.crt -CAkey sample.com.key -set_serial 0 -in www.sample.com.csr -out www.sample.com.crt
# Output
www.sample.com.crt
```
3. Store the cert/key in a Secret Kubernete
```sh
kubectl create secret tls sample-cert --key=www.sample.com.key --cert=www.sample.com.crt
```
4. Update the Gateway to include an HTTPS listener that listens on port 443 and references the _sample-cert_ Secret Kubernete
```sh
kubectl patch gateway eg --type=json --patch '
  - op: add
    path: /spec/listeners/1/tls/certificateRefs/-
    value:
      name: sample-cert
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
        name: sample-cert
      - group: ""
        kind: Secret
        name: sample-cert
      mode: Terminate
```
5. We also update the HTTPRoute _route_serviceb_ to route traffic for hostname _www.sample.com_ to the same backend service
```sh
kubectl patch httproute route-serviceb --type=json --patch '
  - op: add
    path: /spec/hostnames/-
    value: www.sample.com
  '
```
6. Testing
both 2 domain www.example.com and www.sample.com point to the same GATEWAY_HOST ip
```sh
export GATEWAY_HOST=$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}')
curl -v --resolve "www.example.com:443:${GATEWAY_HOST}" --cacert elliptic/envoy/tls/example.com.crt https://www.example.com:443/serviceb
curl -v --resolve "www.sample.com:443:${GATEWAY_HOST}" --cacert elliptic/envoy/tls/sni/sample.com.crt https://www.sample.com:443/serviceb
```