[Guide](https://docs.opensearch.org/latest/data-prepper/common-use-cases/trace-analytics/#trace-tuning)
[Guide Pipeline](https://docs.opensearch.org/latest/data-prepper/pipelines/configuration/sources/otel-trace-source/)
You customize a Data Prepper pipeline to ingest and transform the data for use in OpenSearch. Upon transformation, you can visualize the transformed trace
[Guide OTLP-HTTP](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/otlphttpexporter/README.md)
[Guide Otel Trace Trace Group](https://docs.opensearch.org/latest/data-prepper/pipelines/configuration/processors/otel-trace-group/)

#### Trace Analysis Flow
![Trace Analysis Flow](https://docs.opensearch.org/latest/images/data-prepper/trace-analytics/trace-analytics-components.jpg)

To monitor trace analytics, you need to set up the following components in your service environment:
- Add instrumentation to your application so it can generate telemetry data and send it to an OpenTelemetry collector
- You should configure the collector to export trace data to Data Prepper
- Deploy Data Prepper as the ingestion collector for Open Search
- Use OpenSearch Dashboards to visualize and detect problems in your distributed applications

#### Trace analytics pipeline
![Trace analytics pipeline](https://docs.opensearch.org/latest/images/data-prepper/trace-analytics/trace-analytics-pipeline.jpg)

Processor
There are three processors for the trace analytics feature:
- otel_traces - The otel_traces processor receives a collection of span records from otel-trace-source, and perform stateful processing, extraction, and completion of trace-group-related fields
- otel_traces_group - The otel_traces_group processor fills in the missing trace-group-related fields in the collection of span records by looking up the OpenSearch backend
- service_map - The service_map processor performs the required preprocessing for trace data and builds metadata to display the service-map dashboards

#### Testing with the request from otel-collector
Run the following command to start telemetrygen, which generates synthetic OpenTelemetry traces for 30 seconds and send them to otel-collector:4317 over plaintext gRPC:
```sh
docker run --rm --network=trace-analytics_default \
     ghcr.io/open-telemetry/opentelemetry-collector-contrib/telemetrygen:latest \
     traces \
     --otlp-endpoint=otel-collector:4317 \
     --otlp-insecure \
     --duration=30s \
     --rate=50
```
OR
```sh
~/go/bin/telemetrygen traces --otlp-insecure --traces 1
```

#### Testing with the direct request to data-prepper pipeline
```sh
cat > /tmp/otel-trace3.json <<'JSON'
{
  "resourceSpans": [{
    "resource": {"attributes":[
      {"key":"service.name","value":{"stringValue":"billing"}},
      {"key":"service.version","value":{"stringValue":"2.1.0"}}
    ]},
    "scopeSpans": [{
      "scope": {"name":"manual-https"},
      "spans": [{
        "traceId": "1234567890abcdef1234567890abcdef",
        "spanId":  "feedfacecafebeef",
        "name": "PUT /invoice/42",
        "startTimeUnixNano": "1739999999000000000",
        "endTimeUnixNano":   "1740000000000000000",
        "attributes": [
          {"key":"region","value":{"stringValue":"eu-west-1"}},
          {"key":"retry.count","value":{"intValue":"1"}}
        ]
      }]
    }]
  }]
}
JSON

gzip -c /tmp/otel-trace3.json > /tmp/otel-trace3.json.gz

curl -s -X POST "https://localhost:21890/ingest/otel-trace-pipeline/v1/traces" \
  -H 'Content-Type: application/json' \
  -H 'Content-Encoding: gzip' \
  --insecure \
  --data-binary @/tmp/otel-trace3.json.gz
```

#### Explore
```sh
curl opensearch:9200/_template/otel-v1-apm-service-map-index-template
curl opensearch:9200/_template/otel-v1-apm-span-index-template
```
#### Lookup
```sh
% curl -insecure -ku admin:Skype@123 -XGET https:localhost:9200/_cat/indices
green  open .plugins-ml-config             GP7RMIUeRzyJiJtb0Bk8eg 1 0  1  0    4kb    4kb
green  open top_queries-2026.01.24-65795   AFd0jv2tQ9WJJOtKdb7Few 1 0 17 13 91.9kb 91.9kb
yellow open .opendistro-job-scheduler-lock ZouCQxF-SrK1I8ntxbDeUQ 1 1  1  1 11.5kb 11.5kb
yellow open otel-v1-apm-span-000001        ucPuQ1hRQCaHfOy4W0AVig 1 1  1  0 13.2kb 13.2kb
yellow open otel-v1-apm-service-map      DHhh-oNgQdCmsr1iCKeDIg 1 1  0  0    208b    208b
% curl -insecure -ku admin:Skype@123 -XGET https:localhost:9200/otel-v1-apm-span-000001/_search
% curl -insecure -ku admin:Skype@123 -XGET https://localhost:9200/otel-v1-apm-span-000001/_count 
% curl -insecure -ku admin:Skype@123 -XGET https: localhost:9200/otel-v1-apm-service-map/_search
{"took":3,"timed_out":false,"_shards":{"total":1,"successful":1,"skipped":0,"failed":0},"hits":{"total":{"value":0,"relation":"eq"},"max_score":null,"hits":[]}}
```

