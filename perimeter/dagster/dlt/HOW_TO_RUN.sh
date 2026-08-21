#! /bin/sh
uv venv
source .venv/bin/activate
uv add dlt duckdb
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