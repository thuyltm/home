#### Test with cacert
```sh
docker run -it --network=simple_default -v $(pwd)/otel-root-ca.pem:/home/otel-root-ca.pem   nicolaka/netshoot /bin/sh
grpcurl -cacert /home/otel-root-ca.pem  -d @ data-prepper:21890 opentelemetry.proto.collector.trace.v1.TraceService/Export <<EOF
{
  "resourceSpans": [
    {
      "resource": {
        "attributes": [
          {
            "key": "service.name",
            "value": { "stringValue": "my-trace-service" }
          }
        ]
      },
      "scopeSpans": [
        {
          "spans": [
            {
              "traceId": "0123456789abcdef0123456789abcdef",
              "spanId": "0123456789abcdef",
              "name": "example-span",
              "startTimeUnixNano": "1664827200000000000",
              "endTimeUnixNano": "1664827200001000000",
              "kind": "SPAN_KIND_SERVER"
            }
          ]
        }
      ]
    }
  ]
}
EOF
```

# Test insecure
```sh
docker run -it --network=simple_default nicolaka/netshoot /bin/sh

grpcurl -insecure -d @ data-prepper:21890 opentelemetry.proto.collector.trace.v1.TraceService/Export <<EOF
{
  "resourceSpans": [
    {
      "resource": {
        "attributes": [
          {
            "key": "service.name",
            "value": { "stringValue": "my-trace-service" }
          }
        ]
      },
      "scopeSpans": [
        {
          "spans": [
            {
              "traceId": "0123456789abcdef0123456789abcdef",
              "spanId": "0123456789abcdef",
              "name": "example-span",
              "startTimeUnixNano": "1664827200000000000",
              "endTimeUnixNano": "1664827200001000000",
              "kind": "SPAN_KIND_SERVER"
            }
          ]
        }
      ]
    }
  ]
}
EOF
```

##### Generate CA certificate, public key, private key 


```sh
#!/bin/sh
# Root CA
openssl genrsa -out otel-root-ca-key.pem 2048
openssl req -new -x509 -sha256 -key otel-root-ca-key.pem -subj "/C=VN/ST=HCM/L=VT/O=HOME/OU=IT Dept/CN=data-prepper" -out otel-root-ca.pem -days 730
# Data-Prepper
openssl genrsa -out data-prepper-key-temp.pem 2048
openssl pkcs8 -inform PEM -outform PEM -in data-prepper-key-temp.pem -topk8 -nocrypt -v1 PBE-SHA1-3DES -out data-prepper-key.pem
openssl req -new -key data-prepper-key.pem -subj "/C=VN/ST=HCM/L=VT/O=HOME/OU=IT Dept/CN=data-prepper" -out data-prepper.csr #create a certificate signing request
echo 'subjectAltName=DNS:data-prepper' > data-prepper.ext
openssl x509 -req -in data-prepper.csr -CA otel-root-ca.pem -CAkey otel-root-ca-key.pem -CAcreateserial -sha256 -out data-prepper.pem -days 730 -extfile data-prepper.ext
```

#### Explore grpc service list 
```sh
grpcurl -insecure data-prepper:21890 list
grpcurl -insecure data-prepper:21890 list opentelemetry.proto.collector.trace.v1.TraceService
grpcurl -insecure data-prepper:21890 describe opentelemetry.proto.collector.trace.v1.TraceService.Export
grpcurl -insecure data-prepper:21890 describe .opentelemetry.proto.collector.trace.v1.ExportTraceServiceRequest
```

#### Explore Opensearch indices
```sh
% curl  -XGET http://localhost:9200/_cat/indices
green  open .opensearch-observability      Ok-nsnmuT16P_HJns_qhvw 1 0  0 0    208b    208b
green  open .plugins-ml-config             TtZWDc_1Qr-LmBGdwNL-Jw 1 0  1 0     4kb     4kb
green  open .ql-datasources                ntd2i0puSwi6XskImWDlqQ 1 0  0 0    208b    208b
green  open .opendistro-job-scheduler-lock HAoAxXNQSi24zcVIqHBn_A 1 0  1 1  36.1kb  36.1kb
green  open top_queries-2026.01.24-65795   wVs_IiZqS4uZcRDHIWoEyQ 1 0 32 0 108.1kb 108.1kb
green  open .kibana_1                      OwnVyzQdQDejma_15HQIug 1 0  3 0    24kb    24kb
yellow open my-index-name                  ivxGp_k4RvWX3xhI2qa98w 1 1  2 0  25.5kb  25.5kb
yellow open otel-v1-apm-span-000001        Es0vhGumRQKeKmIj6yL-1g 1 1  0 0    208b    208b

% curl -XGET http://localhost:9200/my-index-name/_mapping
```