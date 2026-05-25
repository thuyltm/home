#! /bin/bash
docker compose up --build
opentelemetry-instrument \
    --traces_exporter otlp \
    --metrics_exporter otlp \
    --logs_exporter otlp \
    --service_name dice-server \
   --exporter_otlp_metrics_insecure true \
  --exporter_otlp_metrics_endpoint  "http://localhost:4317" \
    flask  --app rolldice_instrumentation_metric run