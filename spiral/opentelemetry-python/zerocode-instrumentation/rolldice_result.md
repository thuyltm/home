### Installation opentelemetry libraries
```sh
pip install opentelemetry-distro
opentelemetry-boostrap -a install
```
### Testing
#### Export log information to the console
1. Start Flask server
```sh
opentelemetry-instrument \
    --traces_exporter console \
    --metrics_exporter console \
    --logs_exporter console \
    --service_name dice-server \
    flask --app rolldice run
```
2. Open http://localhost:5000/rolldice in your web browser and reload the page a few times

#### Print logs to the OTLP
1. Install the exporter library
```sh
pip install opentelemetry-exporter-otlp
```

2. Start OTLP Server first
```sh
docker run -p 4317:4317 -v ./otel-collector-config.yaml:/etc/otel-collector-config.yaml otel/opentelemetry-collector:latest --config=/etc/otel-collector-config.yaml
```

3. Start Flask server
```sh
opentelemetry-instrument \
    --traces_exporter otlp \
    --logs_exporter otlp \
    --service_name dice-server \
    flask --app rolldice run
# * Serving Flask app 'app'
# * Debug mode: off
# INFO:werkzeug:WARNING: This is a development server. Do not use it in a production deployment. Use a production WSGI server instead.
# * Running on http://127.0.0.1:5000
# INFO:werkzeug:Press CTRL+C to quit
# WARNING:opentelemetry.exporter.otlp.proto.grpc.exporter:Transient error StatusCode.UNAVAILABLE encountered while exporting logs to localhost:4317, retrying in 0.89s.
# ERROR:opentelemetry.exporter.otlp.proto.grpc.exporter:Failed to export logs to localhost:4317, error code: StatusCode.UNAVAILABLE
# WARNING:opentelemetry.exporter.otlp.proto.grpc.exporter:Transient error StatusCode.UNAVAILABLE encountered while exporting logs to localhost:4317, retrying in 0.92s.
# WARNING:opentelemetry-python.app:Anonymous player is rolling the dice: 2
# INFO:werkzeug:127.0.0.1 - - [23/May/2026 10:49:36] "GET /rolldice HTTP/1.1" 200 -
```

3. Open http://localhost:5000/rolldice in your web browser and reload the page a few times

The logs are appearing on the OTLP Server
```sh
[23/May/2026 11:14:21] "GET /rolldice HTTP/1.1" 200
Trace ID: 
Span ID: 
Flags: 0
        {"resource": {"service.instance.id": "4b0cb7b3-8600-4790-b942-df4d6333681c", "service.name": "otelcol", "service.version": "0.152.1"}, "otelcol.component.id": "debug", "otelcol.component.kind": "exporter", "otelcol.signal": "logs"}
2026-05-23T04:14:25.541Z        info    Traces  {"resource": {"service.instance.id": "4b0cb7b3-8600-4790-b942-df4d6333681c", "service.name": "otelcol", "service.version": "0.152.1"}, "otelcol.component.id": "debug", "otelcol.component.kind": "exporter", "otelcol.signal": "traces", "resource spans": 1, "spans": 1}
2026-05-23T04:14:25.542Z        info    ResourceSpans #0
Resource SchemaURL: 
Resource attributes:
     -> telemetry.sdk.language: Str(python)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.42.1)
     -> service.name: Str(dice-server)
     -> telemetry.auto.version: Str(0.63b1)
ScopeSpans #0
ScopeSpans SchemaURL: https://opentelemetry.io/schemas/1.11.0
InstrumentationScope opentelemetry.instrumentation.flask 0.63b1
Span #0
    Trace ID       : d1b36d2df5bfbdd4fa5d13e4b30e46d9
    Parent ID      : 
    ID             : b9a766ecaeaabe76
    Name           : GET /rolldice
    Kind           : Server
    Start time     : 2026-05-23 04:14:21.179708341 +0000 UTC
    End time       : 2026-05-23 04:14:21.181885514 +0000 UTC
    Status code    : Unset
    Status message : 
    DroppedAttributesCount: 0
    DroppedEventsCount: 0
    DroppedLinksCount: 0
Attributes:
     -> http.method: Str(GET)
     -> http.server_name: Str(127.0.0.1)
     -> http.scheme: Str(http)
     -> net.host.name: Str(127.0.0.1:5000)
     -> http.host: Str(127.0.0.1:5000)
     -> net.host.port: Int(5000)
     -> http.target: Str(/rolldice)
     -> net.peer.ip: Str(127.0.0.1)
     -> net.peer.port: Int(51806)
     -> http.user_agent: Str(Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:150.0) Gecko/20100101 Firefox/150.0)
     -> http.flavor: Str(1.1)
     -> http.route: Str(/rolldice)
     -> http.status_code: Int(200)
```