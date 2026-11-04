from dagster_gcp import GCSResource, BigQueryResource
import pandas as pd
from dagster import asset

@asset
def my_gcs_asset(gcs: GCSResource):
    df = pd.DataFrame({"column1": [1, 2, 3], "column2": ["A", "B", "C"]})
    csv_data = df.to_csv(index=False)
    gcs_client = gcs.get_client()
    bucket = gcs_client.bucket("my-city-vt-bucket")
    blob = bucket.blob("my_dataframe.csv")
    blob.upload_from_string(csv_data)

@asset
def my_table(bigquery: BigQueryResource):
    with bigquery.get_client() as client:
        client.query("SELECT * FROM my_dataset.my_table")

@asset
def users_raw() -> list:
    return [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]
