#! /bin/sh
# create the root certificate
openssl req -x509 \
    -new -sha256 -newkey rsa:2048 -nodes \
    -days 3650 -subj '/CN=root' \
    -config example.cnf \
    -extensions ca_exts \
    -keyout root.key -out root.cert
# create the signing request for the intermediate certificate
openssl req \
  -new -sha256 -newkey rsa:2048 -nodes \
  -subj '/CN=int' \
  -config example.cnf \
  -extensions ca_exts \
  -keyout int.key -out int.req
# create the intermediate certificate signed by the ROOT AUTHORITY
openssl x509 -req \
  -set_serial 1 \
  -days 3650 \
  -extfile example.cnf \
  -extensions ca_exts \
  -CAkey root.key -CA root.cert \
  -in int.req -out int.cert

# create the signing request for leaf certificate
openssl req -new -sha256 -newkey rsa:2048 -nodes \
  -subj '/CN=example' \
  -config example.cnf \
  -extensions leaf_exts \
  -keyout leaf.key -out leaf.req
# create the leaf certificate issued by the INTERMEDIATE AUTHORITY
openssl x509 -req \
  -set_serial 1 -days 3650 \
  -extfile example.cnf \
  -extensions leaf_exts \
  -CAkey int.key -CA int.cert \
  -in leaf.req -out leaf.cert


