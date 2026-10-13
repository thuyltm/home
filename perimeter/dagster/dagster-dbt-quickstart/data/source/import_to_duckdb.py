import duckdb

con = duckdb.connect("/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/staging/data.duckdb")

# Automatically create a table and infer schema directly from the CSV
con.execute("CREATE TABLE IF NOT EXISTS trips AS SELECT *  FROM '/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/source/taxi_trips.parquet'")

# Verify the data was imported
con.sql("SELECT * FROM trips LIMIT 5").show()

# Automatically create a table and infer schema directly from the CSV
con.execute("CREATE TABLE IF NOT EXISTS zones AS SELECT * FROM '/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/source/taxi_zones.csv'")

# Verify the data was imported
con.sql("SELECT * FROM zones LIMIT 5").show()
