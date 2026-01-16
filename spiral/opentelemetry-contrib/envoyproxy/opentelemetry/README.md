[Guide](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/opentelemetry)

For _service-1_, requests are routed based on the request path _service/1_, as follows:

User -> Envoy (envoy-front-proxy)->Envoy(envoy-1)->service-1

For _service-2_, requests are routed based on the request path _service/2_, as follows:

User -> Envoy (envoy-front-proxy) -> Envoy(envoy-1)->Envoy(envoy-2)->service-2

#### VIew the traces in OpenTelemetry UI
Point your browser to http://localhost:55679/debug/tracez