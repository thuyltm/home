from pathlib import Path
import dagster as dg
from dagster_dbt import DbtCliResource
from dagster_duckdb import DuckDBResource
from dagster_dbt_quickstart.defs.assets.dbt import dbt_analytics

database_resource_value = dg.EnvVar("DUCKDB_DATABASE").get_value()

database_resource = DuckDBResource(
    database=database_resource_value if database_resource_value is not None else "data/staging/data.duckdb"
)

dbt_resource_value = Path(__file__).joinpath("../..", "analytics").resolve()

dbt_resource = DbtCliResource(
    project_dir=dbt_resource_value if dbt_resource_value is not None else "src/dagster_dbt_quickstart/analytics",
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