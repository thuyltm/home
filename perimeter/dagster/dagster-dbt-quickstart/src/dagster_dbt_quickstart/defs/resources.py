import dagster as dg
from dagster_dbt import DbtCliResource
from dagster_duckdb import DuckDBResource
from dagster_dbt_quickstart.defs.assets.dbt import dbt_analytics


database_resource = DuckDBResource(
    database="/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/staging/data.duckdb"
)

dbt_resource = DbtCliResource(
    project_dir="/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/src/dagster_dbt_quickstart/analytics",
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