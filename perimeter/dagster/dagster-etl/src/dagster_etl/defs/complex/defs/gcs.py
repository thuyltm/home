from dagster_gcp import GCSResource, BigQueryResource
from google.cloud import bigquery as bq
import pandas as pd
import dagster as dg

iris_harvest_data = dg.AssetSpec(key="iris_harvest_data")

############################################################
# Option 1: Using the BigQuery resource
############################################################
@dg.asset
def iris_data(bigquery: BigQueryResource) -> None:
    iris_df = pd.read_csv(
        "https://docs.dagster.io/assets/iris.csv",
        names=[
            "sepal_length_cm",
            "sepal_width_cm",
            "petal_length_cm",
            "petal_width_cm",
            "species",
        ],
    )
    with bigquery.get_client() as client:
        job = client.load_table_from_dataframe(
            dataframe=iris_df,
            destination="iris.iris_data",
        )
        job.result()

# downstream: create a new table that only contains the Iris-setosa species
@dg.asset(deps=[iris_data])
def iris_setosa(bigquery: BigQueryResource) -> None:
    job_config = bq.QueryJobConfig()
    # Overwrite the table if it already exists
    job_config.write_disposition = bq.WriteDisposition.WRITE_TRUNCATE
    # Set your destination table
    job_config.destination = "articulate-run-306102.iris.iris_setosa"
    sql = "SELECT * FROM articulate-run-306102.iris.iris_data WHERE species = 'Iris-setosa'"
    with bigquery.get_client() as client:
        job = client.query(sql, job_config=job_config)
        job.result()

##############################################################
# Google Cloud Storage
##############################################################
@dg.asset
def my_gcs_asset(gcs: GCSResource):
    client = gcs.get_client()
    bucket = client.bucket("dagster123")
    blob = bucket.blob("gcs.py")
    blob.upload_from_string("Hello from Dagster!")


