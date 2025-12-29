#### Collector Components
![Collector Components](pipeline.png)
[Guideline](https://betterstack.com/community/guides/observability/opentelemetry-collector/)
Here is a quick overview of the basic structure of a Collector configuration file
```sh
receivers:
  otlp:
    protocols:
      http:
        endpoint: 0.0.0.0:4318
processors:
  batch:
exporters:
  otlp:
    endpoint: jaeger:4317
extensions:
  health_check:
service:
  extensions: [health_check]
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp]
```
This configuration sets up an OpenTelemetry Collector that receives trace data via the OTLP protocol over HTTP on port 4318, applies batch processing, and then exports the processed traces to a Jaeger endpoint located at jaeger:4317. It also includes a health_check extension for monitoring the collector's status
### Receiver
Receivers are the components responsible for collecting telemetry data from various sources, serving as the entry points into the Collector. They gather traces, metrics, and logs from instrumented applications, agents, or other systems, and translate the incoming data into OpenTelemetry's internal format, preparing it for further processing and export

The contrib repository boasts over 90 additional receivers, catering to a wide array of data formats and protocols, including popular sources like Kafka, PostgresSQL, Redis, GCP Pubsub [Reference](https://opentelemetry.io/docs/collector/components/receiver/)
```sh
receivers:
  postgresql:
    endpoint: postgresql:5432
    username: root
    password: otel
    metrics:
      postgresql.blks_hit:
        enabled: true
      postgresql.blks_read:
        enabled: true
      postgresql.tup_fetched:
        enabled: true
      postgresql.tup_returned:
        enabled: true
      postgresql.tup_inserted:
        enabled: true
      postgresql.tup_updated:
        enabled: true
      postgresql.tup_deleted:
        enabled: true
      postgresql.deadlocks:
        enabled: true
    tls:
      insecure: true
  redis:
    endpoint: "valkey-cart:6379"
    username: "valkey"
    collection_interval: 10s
```
### Processors
Processors are components that modify or enhance telemetry data as it flows through the pipeline. They perform various operations on the collected telemetry data, such as filtering, transforming, enriching, and batching so that it is ready to be exported

This processor groups spans, metrics, or logs into time-based and size-based batches, enhancing efficiency. Additionaly, it supports sharding data based on client metadata, allowing for effective multi-tenant data processing even with high volumes.

Beyond these, the contrib repository offers several other processors for tasks like filtering sensitive data, appending Kubernetes metadata [Reference](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor)

### Exporters
Exporters serve as the final stage in the Collector's pipeline and are responsible for sending processed telemetry data to various backend systems such as observability platform, databases, or cloud services, where the data is soted, visualized, and analyzed

### Extensions
Extensions offer features like health checks, performance profiling, authentication and integration with external systems
```sh
extensions:
  pprof:
  health_check:
  zpages:
```
The pprof extension here enables Go's net/http/pprof endpoint on http://localhost:1777 so that you can collect performance profiles and investigate issues with the service

The health_check extension offers an HTTP URL (http://localhost:13133/ by default) that can be used to monitor the collector's status

The zPages extension provides various HTTP endpoints (http://localhost:55679/debug) for monitoring and debugging the Collector without relying on any backend. This enables you to inspect traces, metrics, and the collector's internal state directly, assisting in troubleshooting and performance optimization. [Reference](https://github.com/open-telemetry/opentelemetry-collector/tree/main/extension/zpagesextension)

### Connectors
Connectors are specialized components that bridge the different pipelines within the OpenTelemetry Collector. They function as both an exporter for one pipeline and a receiver for another, allowing telemetry data to flow seamlessly between pipelines
```sh
connectors:
  count:
    logs:
      app.event.count:
        description: "Log count by event"
        attributes:
          - key: event
```
For instance, the count connector can count various telemetry data types. It groups incoming logs based on the event attribute and counts the occurrences of each event type. The result is exported as the metric `app.event.count`

### Services
The service section specifies which components, such as receivers, processors, exporters, connectors, and extensions, are active and how they interconnected through pipelines
```sh
service:
  extensions: [health_check, pprof, zpages]
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp]
    metrics:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp]
    logs:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp]
  telemetry:
    metrics:
      level: detailed
      readers:
        - periodic:
            interval: 10000000
            timeout: 5000
            exporter:
              otlp:
                protocol: http/protobuf
                endpoint: "http://otel-collector:4318"
```
Metrics are exposed through a Prometheus interface, which defaults to port 8888

Logs provide insights into Collector events lile startups, shutdowns, data drops, and crashes

Traces add the ability to expose Collector's internal telemetry traces

```