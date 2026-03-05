[Guide](https://gateway.envoyproxy.io/docs/tasks/observability/proxy-trace/)
#### Install the Gateway API CRDs and Envoy Gateway using Helm:
```sh
helm install eg oci://docker.io/envoyproxy/gateway-helm --version v1.6.2 -n envoy-gateway-system --create-namespace
```
#### Install Add-ons
Envoy Gateway provides an add-ons Helm chart to simplify the installation of observability components.

By default, the OpenTelemetry Collector is disabled. To install add-ons with OpenTelemetry Collector enabled, use the following command
```sh
helm install eg-addons oci://docker.io/envoyproxy/gateway-addons-helm --version v1.6.1 --set opentelemetry-collector.enabled=true -n monitoring --create-namespace
```

By default, Envoy Gateway doesn't send traces to any sink. You can enable traces by setting the _telemetry.tracing_ in the EnvoyProxy CRD