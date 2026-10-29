#! /bin/sh
uv run dagster api grpc -h 0.0.0.0 -p 4000 -m dagster_postgresql_quickstart.definitions

source .venv/bin/activate

# Build the container image
docker build -t thuyltm2201/pipes-example:latest .
docker push docker.io/thuyltm2201/pipes-example:latest