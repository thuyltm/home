1. Start OpenTelemetry Collector
```sh
docker run -p 4317:4317 -p 4318:4318 --rm -v $(pwd)/otel-collector-config.yml:/etc/otel-collector-config.yml otel/opentelemetry-collector:latest --config=/etc/otel-collector-config.yml
```
2. Start Flask Server
```sh
flask --app app run
```
3. Nagivate to **http://localhost:5000/rolldice** in your web browser
4. OpenTelemetry Collector output is shown below:
```sh
#info    Traces  {"resource": {"service.instance.id": "51c9adb7-d1e3-49e6-adea-189cd108d1e7", "service.name": "otelcol", "service.version": "0.152.1"}, "otelcol.component.id": "debug", "otelcol.component.kind": "exporter", "otelcol.signal": "traces", "resource spans": 1, "spans": 1}
#info    ResourceSpans #0
#Resource SchemaURL: 
#Resource attributes:
#     -> telemetry.sdk.language: Str(python)
#     -> telemetry.sdk.name: Str(opentelemetry)
#     -> telemetry.sdk.version: Str(1.42.1)
#     -> service.name: Str(flask-test)
#ScopeSpans #0
#ScopeSpans SchemaURL: 
#InstrumentationScope diceroller.tracer 
#Span #0
#    Trace ID       : 9cd163c6c1ea95439a7d36b1ed853812
#    Parent ID      : 
#    ID             : 29465994e51d5a7f
#    Name           : roll
#    Kind           : Internal
#    Start time     : 2026-05-25 08:50:59.075210562 +0000 UTC
#    End time       : 2026-05-25 08:50:59.076342551 +0000 UTC
#    Status code    : Unset
#    Status message : 
#    DroppedAttributesCount: 0
#    DroppedEventsCount: 0
#    DroppedLinksCount: 0
#Attributes:
#     -> roll.value: Str(6)
#        {"resource": {"service.instance.id": "51c9adb7-d1e3-49e6-adea-189cd108d1e7", "service.name": "otelcol", "service.version": "0.152.1"}, "otelcol.component.id": "debug", "otelcol.component.kind": "exporter", "otelcol.signal": "traces"}
```
Flash Server Output is as follow:
```sh
#{
#    "body": "WARNING:opentelemetry-python.app:Anonymous player is rolling the dice: 6",
#    "severity_number": 13,
#    "severity_text": "WARN",
#    "attributes": {
#        "code.file.path": "/home/thuy/Documents/Learn/home/spiral/opentelemetry-python/app.py",
#        "code.function.name": "roll_dice",
#        "code.line.number": 76
#    },
#    "dropped_attributes": 0,
#    "timestamp": "2026-05-25T08:50:59.075936Z",
#    "observed_timestamp": "2026-05-25T08:50:59.076154Z",
#    "trace_id": "0x9cd163c6c1ea95439a7d36b1ed853812",
#    "span_id": "0x29465994e51d5a7f",
#    "trace_flags": 3,
#    "resource": {
#        "attributes": {
#            "telemetry.sdk.language": "python",
#            "telemetry.sdk.name": "opentelemetry",
#            "telemetry.sdk.version": "1.42.1",
#            "service.name": "unknown_service"
#        },
#        "schema_url": ""
#    },
#    "event_name": ""
#}
#{
#    "body": "INFO:werkzeug:127.0.0.1 - - [25/May/2026 15:50:59] \"GET /rolldice HTTP/1.1\" 200 -",
#    "severity_number": 9,
#    "severity_text": "INFO",
#    "attributes": {
#        "code.file.path": "/home/thuy/miniconda3/envs/myenv/lib/python3.14/site-packages/werkzeug/_internal.py",
#        "code.function.name": "_log",
#        "code.line.number": 97
#    },
#    "dropped_attributes": 0,
#    "timestamp": "2026-05-25T08:50:59.077121Z",
#    "observed_timestamp": "2026-05-25T08:50:59.077181Z",
#    "trace_id": "0x00000000000000000000000000000000",
#    "span_id": "0x0000000000000000",
#    "trace_flags": 0,
#    "resource": {
#        "attributes": {
#            "telemetry.sdk.language": "python",
#            "telemetry.sdk.name": "opentelemetry",
#            "telemetry.sdk.version": "1.42.1",
#            "service.name": "unknown_service"
#        },
#        "schema_url": ""
#    },
#    "event_name": ""
#}
```