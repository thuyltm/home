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
##################################################
# pip install dagster-aws
pip install dagster-gcp
############################################################
# Confirm the assets and resources are configured correctly
############################################################
dg check defs
############################################################
# Start the Dagster webserver
############################################################
dg dev
############################################################
# Prepare dataset, table
############################################################
bq mk --location=US --dataset articulate-run-306102:iris
bq mk --table articulate-run-306102:iris.iris_data
bq mk --table articulate-run-306102:iris.iris_setosa

# Use this link https://console.cloud.google.com/iam-admin/iam?project=articulate-run-306102 
# to assign the required Bigquery role (bigquery.jobUser and bigquery.dataEditor) for account thuy-le-thi-minh@articulate-run-306102.iam.gserviceaccount.com
############################################################
# View content of database Duckdb
############################################################
duckdb data.duckdb
D SHOW TABLES;
D SELECT * FROM raw_asteroid_data;
############################################################
# Integrate dlt with dagster
############################################################
uv add dagster dagster-dlt dlt
python dlt_quickstart.py
duckdb data.duckdb
D DESCRIBE;
D SELECT schema_name, database_name, internal 
  FROM duckdb_schemas();
D SELECT table_name
  FROM information_schema.tables
  WHERE table_schema = 'mydata';
D DESCRIBE mydata.load_dict;
D SELECT * FROM mydata.load_dict;
D SELECT * FROM nasa_neo.fetch_neo_data;
#######################################################################
# Incorporate dlt with dagster
# Sling is a lightweight and hight performance ETL data movement tool
#######################################################################
