import dagster as dg
from dagster_duckdb import DuckDBResource

@dg.definitions
def resources():
    return dg.Definitions(
        resources={
            "complexDatabase": DuckDBResource(
                database="src/dagster_etl/data/staging/data.duckdb",
            ),
        }
    )