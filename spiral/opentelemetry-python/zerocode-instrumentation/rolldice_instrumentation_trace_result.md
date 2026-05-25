Body: Str(127.0.0.1 - - [24/May/2026 09:55:53] "GET /rolldice HTTP/1.1" 200 -)
Trace ID: 
Span ID: 
Flags: 0
 {"resource": {"service.instance.id": "638b1fe7-f362-4fc7-b31b-0f4194e09848", "service.name": "otelcol", "service.version": "0.152.1"}, "otelcol.component.id": "debug", "otelcol.component.kind": "exporter", "otelcol.signal": "logs"}
2026-05-24T02:55:55.766Z info    Traces  {"resource": {"service.instance.id": "638b1fe7-f362-4fc7-b31b-0f4194e09848", **"service.name": "otelcol"**, "service.version": "0.152.1"}, "otelcol.component.id": "debug", "otelcol.component.kind": "exporter",**"otelcol.signal": "traces"**, "resource spans": 1, "spans": 2}
2026-05-24T02:55:55.766Z info    ResourceSpans #0
Resource SchemaURL: 
**Resource attributes**:
     -> telemetry.sdk.language: Str(python)
     -> telemetry.sdk.name: Str(opentelemetry)
     -> telemetry.sdk.version: Str(1.42.1)
     -> **service.name: Str(dice-server)**
     -> telemetry.auto.version: Str(0.63b1)
ScopeSpans #0
ScopeSpans SchemaURL: 
**InstrumentationScope diceroller.tracer**
Span #0
    Trace ID       : 0e6df78f9a0a16f0e0a6de7f09c6a20a
    Parent ID      : 9a9c003f7be5c274
    ID             : 7a710c5129288d83
    Name           : roll
    Kind           : Internal
    Start time     : 2026-05-24 02:55:53.654842858 +0000 UTC
    End time       : 2026-05-24 02:55:53.654868996 +0000 UTC
    Status code    : Unset
    Status message : 
    DroppedAttributesCount: 0
    DroppedEventsCount: 0
    DroppedLinksCount: 0
**Attributes**:
     -> **roll.value: Int(4)**
ScopeSpans #1
ScopeSpans SchemaURL: https://opentelemetry.io/schemas/1.11.0
InstrumentationScope opentelemetry.instrumentation.flask 0.63b1
Span #0
    Trace ID       : 0e6df78f9a0a16f0e0a6de7f09c6a20a
    Parent ID      : 
    ID             : 9a9c003f7be5c274
    Name           : GET /rolldice
    Kind           : Server
    Start time     : 2026-05-24 02:55:53.654095556 +0000 UTC
    End time       : 2026-05-24 02:55:53.655109144 +0000 UTC
    Status code    : Unset
    Status message : 
    DroppedAttributesCount: 0
    DroppedEventsCount: 0
    DroppedLinksCount: 0
Attributes:
     -> http.method: Str(GET)
     -> http.server_name: Str(127.0.0.1)
     -> http.scheme: Str(http)
     -> net.host.name: Str(localhost:5000)
     -> http.host: Str(localhost:5000)
     -> net.host.port: Int(5000)
     -> http.target: Str(/rolldice)
     -> net.peer.ip: Str(127.0.0.1)
     -> net.peer.port: Int(49162)
     -> http.user_agent: Str(Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:150.0) Gecko/20100101 Firefox/150.0)
     -> http.flavor: Str(1.1)
     -> http.route: Str(/rolldice)
     -> http.status_code: Int(200)
 {"resource": {"service.instance.id": "638b1fe7-f362-4fc7-b31b-0f4194e09848", "service.name": "otelcol", "service.version": "0.152.1"}, "otelcol.component.id": "debug", "otelcol.component.kind": "exporter", "otelcol.signal": "traces"}