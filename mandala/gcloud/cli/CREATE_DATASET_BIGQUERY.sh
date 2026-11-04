#! /bin/sh
##################################################
# Create a dataset
#################################################
bq mk \
  --location=US \
  --description="My first BigQuery dataset" \
  -d articulate-run-306102:my_dataset
##################################################
# Create a table
##################################################
bq mk \
  -t articulate-run-306102:my_dataset.my_table \
  id:INTEGER,name:STRING,created_at:TIMESTAMP
#####################################################
# APPEND specific JSON records into an existing table
#####################################################
bq load \
  --source_format=CSV \
  --skip_leading_rows=1 \
  articulate-run-306102:my_dataset.my_table \
  data.csv \
  id:INTEGER,name:STRING,created_at:TIMESTAMP


