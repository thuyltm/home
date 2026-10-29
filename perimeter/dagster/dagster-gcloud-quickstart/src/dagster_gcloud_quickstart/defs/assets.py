from dagster_gcp import GCSResource, BigQueryResource
import pandas as pd
from dagster import Definitions, asset
from dagster_gcloud_quickstart.defs.resources import PostgresResource

@asset
def my_gcs_asset(gcs: GCSResource):
    df = pd.DataFrame({"column1": [1, 2, 3], "column2": ["A", "B", "C"]})
    csv_data = df.to_csv(index=False)
    gcs_client = gcs.get_client()
    bucket = gcs_client.bucket("my-cool-bucket")
    blob = bucket.blob("my_dataframe.csv")
    blob.upload_from_string(csv_data)

@asset
def my_table(bigquery: BigQueryResource):
    with bigquery.get_client() as client:
        client.query("SELECT * FROM my_dataset.my_table")

@asset
def users_raw() -> list:
    return [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]

@asset
def postgres_users_table(postgres: PostgresResource, users_raw: list):
    """An asset that populates a table inside PostgreSQL."""
    # Write logic to insert 'users_raw' into Postgres
    query = "INSERT INTO users (id, name) VALUES (%s, %s)..."
    postgres.execute_query(query)
    
    # Returning metadata so Dagster can track row counts, schema, etc.
    return {
        "metadata": {
            "row_count": len(users_raw),
            "table_name": "users"
        }
    }