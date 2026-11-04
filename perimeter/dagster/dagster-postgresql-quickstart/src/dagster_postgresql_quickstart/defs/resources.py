import dagster as dg
from dagster_k8s import PipesK8sClient
from dagster_postgresql_quickstart.defs.asset import PostgresResource
from dagster_gcp.gcs import GCSResource
from dagster_gcp import BigQueryResource

@dg.definitions
def resources():
    return dg.Definitions(
        resources={
            "k8s_pipes_client": PipesK8sClient(),
            "postgres": PostgresResource(
                connection_string="postgresql://test_user:test_pass@my-postgres.default.svc.cluster.local:5432/test_db"
            ),
            "gcs": GCSResource(project="articulate-run-306102"),
            "bigquery": BigQueryResource(project="articulate-run-306102")
        }
    )