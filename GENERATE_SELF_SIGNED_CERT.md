If you don't have access to a certificate authorization (CA) for your organization and want to use your application in production environment, you can generate your own self-signed certificates using OpenSSL
### Generate a root certificate
Root CA (Root __Certificate Authority__), most trusted foundation entity in a digital trust chain, digitally signs and issues certificates that allow applications to be securely verified on the internet
```sh
# Generate a Root CA private key
openssl genrsa -out root-ca-key.pem 2048
# Generate a root certificate
openssl req -new -x509 -sha256 -key root-ca-key.pem -out root-ca.pem -days 730 \
  -addext 'basicConstraints = critical, CA:TRUE, pathlen:0' \
  -addext 'keyUsage = critical, keyCertSign, cRLSign' \
  -addext 'authorityKeyIdentifier = keyid'
```
### Generate a client certificate
1. To generate a client certificate, first create __a client private key__
```sh
openssl genrsa -out client-key.pem 2048
```
2. create __a certificate signing request__ (CSR)
```sh
openssl req -new -key client-key.pem -subj "/C=CA/ST=ONTARIO/L=TORONTO/O=ORG/OU=UNIT/CN=localhost" -out client.csr
```
3. Generate __the client certificate__ by submiting CSR to CA root for signing
```sh
openssl x509 -req -in client.csr -CA root-ca.pem -CAkey root-ca-key.pem -CAcreateserial -sha256 -out client.pem -days 730
```

Convert PEM certificates to keystore and truststore files
- Keystore: The keystore hold private, sensitive data used to authenticate your application's identity to external entities
- Truststore: The truststore holds all verified external entities