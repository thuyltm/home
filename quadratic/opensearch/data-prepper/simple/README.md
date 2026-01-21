[Guide](https://github.com/opensearch-project/data-prepper/blob/main/docs/simple_pipelines.md)

[Guide](https://docs.opensearch.org/latest/security/configuration/generate-certificates/)

To generate self-signed certificates, 
1. the first step in this process is to generate a private key using the _openssl genrsa_ commands. As the name suggests, you should keep this file private. 
2. So as to generate a root certificate, you use the private key to generate a self-signed certificate for the root CA
```sh
# 1. Generate a private key (e.g., 4096-bit RSA or an ECC key)
openssl genrsa -aes256 -out rootCA.key 4096
# OR for ECC:
# openssl ecparam -genkey -name secp384r1 | openssl ec -aes256 -out rootCA.key

# 2. Create and self-sign the root certificate (valid for 10 years)
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 3650 -out rootCA.crt

# 3. (Optional) Convert to .cer format (Base64 encoded) if needed for specific systems
openssl x509 -in rootCA.crt -outform DER -out rootCA.cer

```
3. 
- Aim to generate an admin certificate, you create a certificate signing request (CSR). This file acts as an request to a CA for a signed certificate. 
- Now that the private key and signing request have been created, generate the certificate

```sh
#!/bin/sh
# Root CA
openssl genrsa -out root-ca-key.pem 2048
openssl req -new -x509 -sha256 -key root-ca-key.pem -subj "/C=VN/ST=HCM/L=VT/O=HOME/OU=IT Dept/CN=localhost" -out root-ca.pem -days 730
# Data-Prepper
openssl genrsa -out data-prepper-key-temp.pem 2048
openssl pkcs8 -inform PEM -outform PEM -in data-prepper-key-temp.pem -topk8 -nocrypt -v1 PBE-SHA1-3DES -out data-prepper-key.pem
openssl req -new -key data-prepper-key.pem -subj "/C=VN/ST=HCM/L=VT/O=HOME/OU=IT Dept/CN=localhost" -out data-prepper.csr #create a certificate signing request
echo 'subjectAltName=DNS:localhost' > data-prepper.ext
openssl x509 -req -in data-prepper.csr -CA root-ca.pem -CAkey root-ca-key.pem -CAcreateserial -sha256 -out data-prepper.pem -days 730 -extfile data-prepper.ext
```

Aim to convert PEM certificates to keystore and trustore files
```sh
#!/bin/sh

# Convert node certificate
cat root-ca.pem data-prepper.pem data-prepper-key.pem > combined-data-prepper.pem
echo "Enter password for private key"
openssl pkcs12 -export -in combined-data-prepper.pem -out data-prepper-cert.p12 -name data-prepper
echo "Enter password for keystore"
keytool -importkeystore -srckeystore data-prepper-cert.p12 -srcstoretype pkcs12 -destkeystore data-prepper.jks

# Import certificates to truststore
keytool -importcert -keystore truststore.jks -file root-ca.cer -storepass changeit -trustcacerts -deststoretype pkcs12

# Cleanup
rm combined-admin.pem
rm combined-node1.pem
```

#### Question Security
Owner: CN=localhost, OU=IT Dept, O=Home, L=VT, ST=HCM, C=VN
Issuer: CN=localhost, OU=IT Dept, O=Home, L=VT, ST=HCM, C=VN
/C=CA/ST=VIETNAM/L=HCM/O=ORG/OU=HCM/CN=localhost

What is your first and last name? localhost
What is the name of your organizational unit? IT Dept
What is the name of your organization? Home
What is the name of your City or Locality? VT
What is the name of your State or Province? HCM
What is the two-letter country code for this unit? VN

#### Check information in p12 file
```sh
keytool -delete -alias localhost -keystore keystore.p12
keytool -list -v -keystore mykeystore.jks 
```

#### Quickly generate keystore (p12 or jks) file
```sh
keytool -genkeypair -alias localhost -keyalg RSA -keysize 2048 -validity 365 -storetype PKCS12 -keystore keystore.p12 -storepass 123456 -keypass 123456
```
 
#### Testing
```sh
curl --cacert root-ca.pem https://localhost:4900/list
```


