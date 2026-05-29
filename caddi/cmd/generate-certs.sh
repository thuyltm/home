#!/bin/sh
# Root CA
openssl genrsa -out root-ca-key.pem 2048
openssl req -new -x509 -sha256 -key root-ca-key.pem -subj "/C=CA/ST=ONTARIO/L=TORONTO/O=ORG/OU=UNIT/CN=root" -out root-ca.pem -days 730 \
  -addext 'basicConstraints = critical, CA:TRUE, pathlen:0' \
  -addext 'keyUsage = critical, keyCertSign, cRLSign' \
  -addext 'authorityKeyIdentifier = keyid'

# WebServer cert
openssl genrsa -out server-key-temp.pem 2048
openssl pkcs8 -inform PEM -outform PEM -in server-key-temp.pem -topk8 -nocrypt -v1 PBE-SHA1-3DES -out server-key.pem
openssl req -new -key server-key.pem -subj "/C=CA/ST=ONTARIO/L=TORONTO/O=ORG/OU=UNIT/CN=caddi.cmd.com" -out server.csr
openssl x509 -req -in server.csr -CA root-ca.pem -CAkey root-ca-key.pem -CAcreateserial -sha256 -out server.pem -days 730