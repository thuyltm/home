import dagster as dg
from dagster_sling import SlingConnectionResource, SlingResource

source = SlingConnectionResource(
    name="MY_POSTGRES",
    type="postgres",
    host="localhost",
    port=5432,
    database="test_db",
    user="test_user",
    password="test_pass",
)

destination = SlingConnectionResource(
    name="MY_DUCKDB",
    type="duckdb",
    connection_string="duckdb:///home/thuy/Documents/Learn/home/perimeter/dagster/dagster-etl/src/dagster_etl/data/staging/migration.db",
)

sling = SlingResource(
    connections=[
        source,
        destination,
    ]
)

@dg.definitions
def resources():
    return dg.Definitions(
        resources={
            "sling": sling,
        },
    )