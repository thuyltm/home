#! /bin/sh
uv run dagster api grpc -h 0.0.0.0 -p 4000 -m dagster_postgresql_quickstart.definitions

source .venv/bin/activate

# Build the container image
docker build -t thuyltm2201/pipes-example:latest .
docker push docker.io/thuyltm2201/pipes-example:latest
# export your pyproject.toml dependencies to a clean, flat list using the --no-hashes flag
uv export --format requirements-txt --no-hashes --output-file requirements.txt
