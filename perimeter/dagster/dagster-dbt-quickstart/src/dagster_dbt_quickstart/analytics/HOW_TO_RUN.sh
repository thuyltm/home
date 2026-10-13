#! /bin/sh
######################################################################################
# Since you are using sources/raw_taxis.yml, creating tables is a prerequisite
# Import data directly into DuckDB using the DuckDB tool or python instead of dbt seed
######################################################################################
# Use Duckdb
CREATE TABLE zones AS SELECT * FROM '/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/source/taxi_zones.csv';
CREATE TABLE trips AS SELECT * FROM '/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/source/taxi_trips.parquet' WHERE tpep_pickup_datetime BETWEEN '2026-01-01 00:00:00' AND '2026-01-01 23:59:59';
# Use python
python import_to_duckdb.py
######################################################################################
# To compile the folders from a source directory into a target directory
######################################################################################
dbt deps
# Running with dbt=1.12.5
# Installing dbt-labs/dbt_utils
# Installed from version 1.1.1
# Updated version available: 1.4.1
#  
# Updates available for packages: ['dbt-labs/dbt_utils']                 
# Update your versions in packages.yml, then run dbt deps
dbt parse
# Running with dbt=1.12.5
# Registered adapter: duckdb=1.11.0
# Unable to do partial parsing because saved manifest not found. Starting full parse.
# Performance info: /home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/src/dagster_dbt_quickstart/analytics/target/perf_info.json
dbt compile --profiles-dir /home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/src/dagster_dbt_quickstart/analytics
# Running with dbt=1.12.5
# Registered adapter: duckdb=1.11.0
# Found 4 models, 2 data tests, 2 sources, 614 macros
#  
# Concurrency: 1 threads (target='dev')
# The dbt parse command is used to generate a manifest.json for a dbt project.
dbt build --profiles-dir /home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/src/dagster_dbt_quickstart/analytics
dbt parse
#######################################################################################
# RETRIEVE
#######################################################################################
select cast(tpep_pickup_datetime as date) as day from trips group by day order by day;

