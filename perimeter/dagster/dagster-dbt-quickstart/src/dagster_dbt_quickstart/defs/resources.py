import dagster as dg
from dagster_dbt import DbtCliResource
from dagster_duckdb import DuckDBResource
from dagster_dbt_quickstart.defs.assets.dbt import dbt_analytics


database_resource = DuckDBResource(
    database=dg.EnvVar("DUCKDB_DATABASE")
)

dbt_resource = DbtCliResource(
    project_dir=dg.EnvVar("DBT_PROJECT"),
)

@dg.definitions
def resources(): 
    return dg.Definitions(
        assets=[dbt_analytics],
        resources={
            "database": database_resource,
            "dbt": dbt_resource,
        },
    )