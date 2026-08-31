#! /bin/sh
#################################################
# Install psql
#################################################
sudo apt update
sudo apt install postgresql postgresql-contrib
#################################################
# Start Postgres Server
#################################################
docker compose up
#############################################################
# Connect to Postgres DB
#############################################################
psql "postgresql://test_user:test_pass@localhost:5432/test_db"
# \dt test_schema.*
# \d+ test_schema.customers
# select * from test_schema.customers;
# select * from test_schema.products;
# select * from test_schema.orders;

