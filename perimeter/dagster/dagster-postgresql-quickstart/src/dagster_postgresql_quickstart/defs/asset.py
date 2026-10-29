import dagster as dg
from dagster_k8s import PipesK8sClient
from dagster import ConfigurableResource
from sqlalchemy import create_engine
import pandas as pd

@dg.asset
def raw_transactions():
    return [1, 2, 3]

daily_partitions = dg.DailyPartitionsDefinition(start_date="2024-01-01")

@dg.asset(partitions_def=daily_partitions)
def k8s_pipes_asset(context: dg.AssetExecutionContext, k8s_pipes_client: PipesK8sClient):
    partition_key = context.partition_key

    # Use the partition key in your logic or queries
    context.log.info(f"Processing data for partition: {partition_key}")
    return k8s_pipes_client.run(
        context=context, image="docker.io/thuyltm2201/pipes-example:latest", extras={"some_parameter":1}
    ).get_materialize_result()

class PostgresResource(ConfigurableResource):
    connection_string: str

    def get_engine(self):
        return create_engine(self.connection_string)

@dg.asset
def my_database_asset(postgres: PostgresResource) -> None:
    # Create sample data or process a DataFrame
    df = pd.DataFrame({"id": [1, 2, 3], "name": ["Alice", "Bob", "Charlie"]})

    # Save to PostgreSQL
    engine = postgres.get_engine()
    df.to_sql("my_table", engine, if_exists="replace", index=False)