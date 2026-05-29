1. OpenSearch includes demo certificates for quick setup and testing, located in the <OPENSEARCH_HOME>/config/ directory. These certificates are not secure for production environments and must be replaced with your own
```sh
% docker exec -it opensearch-node1 sh
$ ls config/
esnode-key.pem      jvm.options    kirk.pem                  opensearch-notifications-core  opensearch-security            opensearch.yml
esnode.pem          jvm.options.d  log4j2.properties         opensearch-observability       opensearch-security-analytics  root-ca.pem
fips_java.security  kirk-key.pem   opensearch-notifications  opensearch-reports-scheduler   opensearch.keystore            securityadmin_demo.sh
```
2. The incorrect command **curl -XGET http://localhost:9200/_cat/indices** will throws an exception
```sh
[2026-01-24T04:53:48,148][WARN ][i.n.h.s.ApplicationProtocolNegotiationHandler] [opensearch-node1] [id: 0x6cac3785, L:/172.21.0.4:9200 ! R:/172.21.0.1:57774] TLS handshake failed:
opensearch-node1       | io.netty.handler.ssl.NotSslRecordException: not an SSL/TLS record: 474554202f5f6361742f696e646963657320485454502f312e310d0a486f73743a206c6f63616c686f73743a393230300d0a557365722d4167656e743a206375726c2f382e31392e302d4445560d0a4163636570743a202a2f2a0d0a0d0a
```
3. The correct command
```sh
% curl -insecure -XGET https://localhost:9200/_cat/indices -u admin:Skype@123
# OR
% docker cp ddc287f8c1df:/usr/share/opensearch/config/root-ca.pem ~/Documents/Learn/home/quadratic/opensearch/secure
# use --cacert to verify the server domain with the trusted certificate
% curl --cacert esnode.pem --resolve "node-0.example.com:9200:127.0.0.1" https://node-0.example.com:9200 -u admin:Skype@123
# will throws "Host name 'opensearch' does not match the certificate subject provided by the peer (CN=node-0.example.com, OU=node, O=node, L=test, C=de)", with the result that you should set hostname alias 'node-0.example.com' for container name 'opensearch'
% curl --cacert esnode.pem --resolve "opensearch:9200:127.0.0.1" -XGET https://opensearch:9200/_cat/indices -u admin:Skype@123
# use --cert and --key to prove your identity to a server
% curl --cert kirk.pem --key kirk-key.pem --cacert esnode.pem --resolve "node-0.example.com:9200:127.0.0.1" https://node-0.example.com:9200
```


### View the content of client certification
```sh
openssl x509 -in esnode.pem -text -noout
#Certificate:
#    Data:
#        Version: 3 (0x2)
#        Serial Number:
#            69:84:a5:11:3d:e7:ce:ca:2d:59:3a:d6:b9:e5:4f:3e:1d:74:c8:b6
#        Signature Algorithm: sha256WithRSAEncryption
#        Issuer: DC=com, DC=example, O=Example Com Inc., OU=Example Com Inc. Root CA, CN=Example Com Inc. Root CA
#        Validity
#            Not Before: Feb 20 17:03:25 2024 GMT
#            Not After : Feb 17 17:03:25 2034 GMT
#        Subject: C=de, L=test, O=node, OU=node, CN=node-0.example.com
#        Subject Public Key Info:
#            Public Key Algorithm: rsaEncryption
#                Public-Key: (2048 bit)
#                Exponent: 65537 (0x10001)
#        X509v3 extensions:
#            X509v3 Subject Alternative Name: 
#                Registered ID:1.2.3.4.5.5, DNS:node-0.example.com, DNS:localhost, IP Address:0:0:0:0:0:0:0:1, IP Address:127.0.0.1
#            X509v3 Key Usage: 
#                Digital Signature, Non Repudiation, Key Encipherment
#            X509v3 Extended Key Usage: 
#                TLS Web Server Authentication, TLS Web Client Authentication
#            X509v3 Basic Constraints: critical
#                CA:FALSE
#            X509v3 Subject Key Identifier: 
#                D3:FA:83:41:A6:35:D2:32:28:C0:28:CB:52:9C:FF:1D:F4:17:CA:DF
#            X509v3 Authority Key Identifier: 
#                17:87:DF:A0:5A:EB:66:12:A7:D5:D0:F8:BA:12:45:3C:B7:2B:00:9C
#    Signature Algorithm: sha256WithRSAEncryption
#    Signature Value: 
openssl x509 -in root-ca.pem -text -noout
#Certificate:
#    Data:
#        Version: 3 (0x2)
#        Serial Number:
#            0d:64:09:99:66:7d:c4:14:ec:41:47:8e:b7:d1:79:61:23:e9:a8:e2
#        Signature Algorithm: sha256WithRSAEncryption
#        Issuer: DC=com, DC=example, O=Example Com Inc., OU=Example Com Inc. Root CA, CN=Example Com Inc. Root CA
#        Validity
#            Not Before: Feb 20 17:00:36 2024 GMT
#            Not After : Feb 17 17:00:36 2034 GMT
#        Subject: DC=com, DC=example, O=Example Com Inc., OU=Example Com Inc. Root CA, CN=Example Com Inc. Root CA
#        Subject Public Key Info:
#            Public Key Algorithm: rsaEncryption
#                Public-Key: (2048 bit)
#                Modulus:
#                Exponent: 65537 (0x10001)
#        X509v3 extensions:
#            X509v3 Basic Constraints: critical
#                CA:TRUE
#            X509v3 Key Usage: critical
#                Digital Signature, Certificate Sign, CRL Sign
#            X509v3 Subject Key Identifier: 
#                17:87:DF:A0:5A:EB:66:12:A7:D5:D0:F8:BA:12:45:3C:B7:2B:00:9C
#            X509v3 Authority Key Identifier: 
#                keyid:17:87:DF:A0:5A:EB:66:12:A7:D5:D0:F8:BA:12:45:3C:B7:2B:00:9C
#                DirName:/DC=com/DC=example/O=Example Com Inc./OU=Example Com Inc. Root CA/CN=Example Com Inc. Root CA
#                serial:0D:64:09:99:66:7D:C4:14:EC:41:47:8E:B7:D1:79:61:23:E9:A8:E2
#    Signature Algorithm: sha256WithRSAEncryption
#    Signature Value:
```

