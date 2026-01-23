#### Test with cacert
```sh
docker run -it --network=simple_default -v $(pwd)/otel-root-ca.pem:/home/otel-root-ca.pem   nicolaka/netshoot /bin/sh
grpcurl -cacert /home/otel-root-ca.pem -d @ data-prepper:21890 opentelemetry.proto.collector.trace.v1.TraceService/Export <<EOF
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
```