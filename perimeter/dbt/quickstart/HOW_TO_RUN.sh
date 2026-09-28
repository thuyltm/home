#! /bin/sh
uv init
uv sync
#######################
# RUN
#######################
set -e
source .venv/bin/activate
#########################
# ADD Dependencies
#########################
uv add dbt-core
uv add duckdb-dbt
####################################################################################
# Quickstart for dbt v1 using DuckDB
####################################################################################
dbt init duckdb_connect
# 1. Firstly, load the CSV file in the seed path into the data warehouse
dbt seed
# 2. Running 3 table models, 4 view models
dbt run
dbt test
dbt docs generate
dbt docs serve
####################################################################################
# CHECKING
####################################################################################
####################################################################################
# DuckDB stores your data in a local .duckdb file on your machine
# The location of this file is defined by the path field in your ~/.dbt/profiles.yml
####################################################################################
ls -lah *.duckdb
duckdb dev.duckdb
dev D. .tables
dev D. SHOW TABLES;
dev D. SHOW ALL TABLES;
dev D. select * from my_first_dbt_model using sample 10 rows;
dev D. SELECT table_schema, table_name 
FROM information_schema.tables 
WHERE table_type = 'VIEW';
