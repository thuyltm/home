from dagster_gcp.gcs import GCSResource
from dagster_gcp import BigQueryResource

import dagster as dg
from dagster import ConfigurableResource
import psycopg2

class PostgresResource(ConfigurableResource):
    connection_string: str
    def execute_query(self, query: str):
        with psycopg2.connect(self.connection_string) as conn:
            with conn.cursor() as cursor:
                cursor.execute(query)

@dg.definitions
def resources() -> dg.Definitions:
    return dg.Definitions(
        resources={
                "gcs": GCSResource(project="articulate-run-306102"),
                "bigquery": BigQueryResource(project="articulate-run-306102")
            }
        )
