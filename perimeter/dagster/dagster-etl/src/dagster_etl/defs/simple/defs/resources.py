import dagster as dg
from dagster_duckdb import DuckDBResource

@dg.definitions
def resources():
    return dg.Definitions(
        resources={
            "simpleDatabase": DuckDBResource(
                database="src/dagster_etl/data/staging/data.duckdb",
            ),
        }
    )