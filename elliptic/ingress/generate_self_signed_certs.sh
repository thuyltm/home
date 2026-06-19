#! /bin/bash

#1. Generate a root certificate
#1.1. Generate a Root CA private key
openssl genrsa -out root-ca-key.pem 2048
#1.2. Generate a root certificate
openssl req -new -x509 -sha256 -key root-ca-key.pem -out root-ca.pem -days 730 \
  -addext 'basicConstraints = critical, CA:TRUE, pathlen:0' \
  -addext 'keyUsage = critical, keyCertSign, cRLSign' \
  -addext 'authorityKeyIdentifier = keyid'


#2. Generate a client certificate for host argocd.example.com
#2.1. To generate a client certificate, first create a client private key
openssl genrsa -out http-key.pem 2048
#2.2. create a certificate signing request (CSR)
openssl req -new -key http-key.pem -subj "/C=CA/ST=ONTARIO/L=TORONTO/O=ORG/OU=UNIT/CN=argocd.example.com" -out http.csr
#2.3 Generate the client certificate by submiting CSR to CA root for signing
openssl x509 -req -in http.csr -CA root-ca.pem -CAkey root-ca-key.pem -CAcreateserial -sha256 -out http.pem -days 730

#3. Generate a client certificate for host grpc.argocd.example.com
#3.1. To generate a client certificate, first create a client private key
openssl genrsa -out grpc-key.pem 2048
#3.2. create a certificate signing request (CSR)
openssl req -new -key grpc-key.pem -subj "/C=CA/ST=ONTARIO/L=TORONTO/O=ORG/OU=UNIT/CN=grpc.argocd.example.com" -out grpc.csr
#3.3 Generate the client certificate by submiting CSR to CA root for signing
openssl x509 -req -in grpc.csr -CA root-ca.pem -CAkey root-ca-key.pem -CAcreateserial -sha256 -out grpc.pem -days 730