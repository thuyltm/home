```sh
docker compose up --build opensearch
```
# Internal Service Communication
Within a default or user-defined Docker network created by Compose, services can automatically discover and communicate with each other using their service names as hostnames
```sh
services:
  otel-collector:
    image: ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-contrib:0.139.0
    networks:
      mynet:
        aliases:
          - otel-collector.local
          - otel-collector.local
  jaeger:
    image: jaegertracing/jaeger:2.11.0
    networks:
      - mynet

networks:
  mynet:
```
In this example, the jaeger service can reach the otel-collector service using http://otel-collector, http://otel-collector.local, or http://otel-collector.local


When you do not specify a network in your compose.yaml file, Docker Compose automatically creates a single, default network for your entire application. All services defined within that file are connected to this network by default, allowing them to communicate with each other using their service names as hostnames.

# How to LoadTest for Service `dice`
They can communicate each other because of using the same default network `opentelemetry-contrib_default`
```sh
docker compose up
docker compose -f ./dice/docker-compose.yml up
docker compose -f locust/docker-compose.yml up
```