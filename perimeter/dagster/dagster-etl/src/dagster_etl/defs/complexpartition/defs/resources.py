import os
from pathlib import Path

import dagster as dg
from dagster_duckdb import DuckDBResource
from dagster_gcp import BigQueryResource, GCSResource

# Set the environment variable directly in Python
gcloud_authen_path = Path(__file__).absolute().parent / f"articulate-run-306102-910c63dbe0d6.json"
os.environ["GOOGLE_APPLICATION_CREDENTIALS"] = str(gcloud_authen_path.resolve())
os.environ["GOOGLE_API_USE_CLIENT_CERTIFICATE"]="False"

@dg.definitions
def resources():
    return dg.Definitions(
        resources={
            "complexDatabase": DuckDBResource(
                database="src/dagster_etl/data/staging/data.duckdb",
            ),
            "bigquery": BigQueryResource(
                project="articulate-run-306102"
            ),
            "gcs": GCSResource(
                 project="articulate-run-306102"
            )
        }
    )