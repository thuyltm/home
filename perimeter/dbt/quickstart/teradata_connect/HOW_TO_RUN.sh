#! /bin/sh
#########################################
# https://trial.teradata.com/dashboard
#########################################
dbt seed
dbt run
dbt test
