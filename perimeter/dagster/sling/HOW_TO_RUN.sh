#! /bin/sh
#########################################################
# Start Postgres DB
#########################################################
docker compose -f docker/db/postgres/docker-compose.yml up
##########################################################
# Sling cli
##########################################################
sling run --pipeline path/to/pipeline.yaml
# one-time copy of a single table
sling run \
  --src-conn "postgresql://test_user:test_pass@localhost:5432/test_db?sslmode=disable" \
  --tgt-conn "duckdb:///home/thuy/Documents/Learn/home/perimeter/dagster/sling/my_duckdb.db" \
  --src-stream "test_schema.customers" \
  --tgt-object "test_schema.customers" \
  --mode full-refresh
sling run -r replication.yaml

