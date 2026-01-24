curl -X PUT -H "Content-Type: application/json" -d '{ "name": "John Doe", "gpa": 3.89, "grad_year": 2022}' \
https://localhost:9200/students/_doc/1 -ku admin:Skype@123


curl -XGET "https://localhost:9200/students/_mapping" -ku admin:Skype@123


curl -XGET "https://localhost:9200/students/_search" -ku admin:Skype@123 -H 'Content-Type: application/json' -d'
{
  "query": {
    "match_all": {}
  }
}
'

curl -X GET "https://localhost:9200/students/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
    "match_all": {}
  }
}'
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
% curl -insecure -XGET https://localhost:9200/_cat/indices -ku admin:Skype@123
# OR
% docker cp ddc287f8c1df:/usr/share/opensearch/config/root-ca.pem ~/Documents/Learn/home/quadratic/opensearch/sercure
% curl --cacert root-ca.pem -XGET https://localhost:9200/_cat/indices -ku admin:Skype@123
```