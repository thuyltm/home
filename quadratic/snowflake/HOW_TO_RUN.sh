#! /bin/sh
uv venv
source .venv/bin/activate
uv add --pre dbt
##########################################################
# Install snowflake cli
##########################################################
uv tool install snowflake-cli
snow -h
snowsql -a QQZQDUR-QS42787 -u thuyltm2201 -o log_level=DEBUG