#! /bin/sh
# scaffold a new dagster project
uvx create-dagster@latest project dagster-etl
# activate the virtual environment
cd dagster-etl
source .venv/bin/activate
# scaffold an assets file
dg scaffold defs dagster.asset assets.py
# install dagster-duckdb library
uv add dagster-duckdb
# scaffold a resources file
dg scaffold defs dagster.resources resources.py
##################################################
# materialize the assets with the following config
##################################################
resources:
  database:
    config:
      database: src/dagster_etl/data/staging/data.duckdb
#"ops": {"op_name": {"config": {"key": "value"}}}}
ops:
  import_file:
    config:
      path: 2018-01-22.csv
###################################################
# check
duckdb src/dagster_etl/data/staging/data.duckdb
D Select * From raw_data;
##################################################
# Install dagster-aws or dagster-gcp
# pip install dagster-aws
pip install dagster-gcp
