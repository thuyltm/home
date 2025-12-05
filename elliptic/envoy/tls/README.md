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

3. A certicate (or end-entity certificate) is a digital document issued to a user, server, or device, which is signed by root certificates
Triple: example.com.crt (CA) + example.com.key (CAkey) + www.example.com.csr (A certificate signing request) = A end-entity certificate
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
curl -v --resolve "www.example.com:8443:127.0.0.1" --cacert elliptic/envoy/tls/example.com.crt https://www.example.com:8443/serviceb
```
Output
```sh
Connected to www.example.com (10.101.19.130) port 443
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
*  CAfile: elliptic/envoy/tls/example.com.crt
*  CApath: /etc/ssl/certs
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384 / X25519 / RSASSA-PSS
* ALPN: server accepted h2
* Server certificate:
*  subject: CN=www.example.com; O=example organization
*  start date: Nov 29 04:46:53 2025 GMT
*  expire date: Nov 29 04:46:53 2026 GMT
*  common name: www.example.com (matched)
*  issuer: O=example Inc.; CN=example.com
*  SSL certificate verify ok.
*   Certificate level 0: Public key type RSA (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type RSA (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* using HTTP/2
* [HTTP/2] [1] OPENED stream for https://www.example.com:443/serviceb
```

6. Test through Gateway
```sh
export GATEWAY_HOST=$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}')
curl -v --resolve "www.example.com:443:${GATEWAY_HOST}" --cacert elliptic/envoy/tls/example.com.crt https://www.example.com:443/serviceb
```

# Advance
1. SSL connection using TLSv1.2 / ECDHE-RSA-CHACHA20-POLY1305 / X25519 / RSASSA-PSS

RSA + ECDSA dual stack certificates

Elliptic Curve Digital Signature Algorithm (ECDSA) offers a variant of the Digital Signature Algorithm (DSA) which uses elliptic-curve cryptography
```sh
export GATEWAY_HOST=$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}')
curl -v "www.example.com:443:${GATEWAY_HOST}" --cacert elliptic/envoy/tls/example.com.crt https://www.example.com:443/serviceb -Isv --ciphers ECDHE-RSA-CHACHA20-POLY1305 --tlsv1.2 --tls-max 1.2
```
Output
```sh
* Connected to www.example.com (10.101.19.130) port 443
* ALPN: curl offers h2,http/1.1
* Cipher selection: ECDHE-RSA-CHACHA20-POLY1305
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
*  CAfile: elliptic/envoy/tls/example.com.crt
*  CApath: /etc/ssl/certs
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-CHACHA20-POLY1305 / X25519 / RSASSA-PSS
* ALPN: server accepted h2
* Server certificate:
*  subject: CN=www.example.com; O=example organization
*  start date: Nov 29 04:46:53 2025 GMT
*  expire date: Nov 29 04:46:53 2026 GMT
*  common name: www.example.com (matched)
*  issuer: O=example Inc.; CN=example.com
*  SSL certificate verify ok.
*   Certificate level 0: Public key type RSA (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
*   Certificate level 1: Public key type RSA (2048/112 Bits/secBits), signed using sha256WithRSAEncryption
* using HTTP/2
* [HTTP/2] [1] OPENED stream for https://www.example.com:443/serviceb
```
2. SNI based Certificate Selection

Server Name Indincation (SNI) is an extension to the Transport Layer Security (TLS) computer networking protocol. The extension allows a server to present one of multiple possible certificates on the same IP address and TPC port number and hence allows multiple secure (HTTPS) websites to be served by the same IP address
Specific example:
```sh
export GATEWAY_HOST=$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}')
curl -v --resolve "www.example.com:443:${GATEWAY_HOST}" --cacert elliptic/envoy/tls/example.com.crt https://www.example.com:443/serviceb
curl -v --resolve "www.sample.com:443:${GATEWAY_HOST}" --cacert elliptic/envoy/tls/sni/sample.com.crt https://www.sample.com:443/serviceb
```