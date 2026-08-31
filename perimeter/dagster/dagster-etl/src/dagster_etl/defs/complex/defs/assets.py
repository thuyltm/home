import csv
from pathlib import Path

import dagster as dg
from dagster_duckdb import DuckDBResource

partitions_def = dg.DailyPartitionsDefinition(
    start_date="2018-01-21",
    end_date="2018-01-24",
)

@dg.asset(
    partitions_def=partitions_def,
)
def import_partition_file(context: dg.AssetExecutionContext) -> str:
    file_path = (
        Path(__file__).absolute().parent / f"../../../data/source/{context.partition_key}.csv"
    )
    return str(file_path.resolve())

@dg.asset(
    kinds={"duckdb"},
    partitions_def=partitions_def,
)
def duckdb_partition_table(
    context: dg.AssetExecutionContext,
    complexDatabase: DuckDBResource,
    import_partition_file,
):
    table_name = "raw_partition_data"
    with complexDatabase.get_connection() as conn:
        table_query = f"""
            create table if not exists {table_name} (
                date date,
                share_price float,
                amount float,
                spend float,
                shift float,
                spread float
            ) 
        """
        conn.execute(table_query)
        conn.execute(
            f"delete from {table_name} where date = '{context.partition_key}';"
        )
        conn.execute(f"copy {table_name} from '{import_partition_file}';")

#################################################################################
# Dynamic Partition
#################################################################################
dynamic_partitions_def = dg.DynamicPartitionsDefinition(name="dynamic_partition")

@dg.asset(
    partitions_def=dynamic_partitions_def,
)
def import_dynamic_partition_file(context: dg.AssetExecutionContext) -> str:
    file_path = (
        Path(__file__).absolute().parent / f"../../../data/source/{context.partition_key}.csv"
    )
    return str(file_path.resolve())

@dg.asset(
    kinds={"duckdb"},
    partitions_def=dynamic_partitions_def,
)
def duckdb_dynamic_partition_table(
    context: dg.AssetExecutionContext,
    complexDatabase: DuckDBResource,
    import_dynamic_partition_file,
):
    table_name = "raw_dynamic_partition_data"
    with complexDatabase.get_connection() as conn:
        table_query = f"""
            create table if not exists {table_name} (
                date date,
                share_price float,
                amount float,
                spend float,
                shift float,
                spread float
            ) 
        """
        conn.execute(table_query)
        conn.execute(
            f"delete from {table_name} where date = '{context.partition_key}';"
        )
        conn.execute(
            f"copy {table_name} from '{import_dynamic_partition_file}';"
        )

# Cloud Storage
class IngestionFileS3Config(dg.Config):
    bucket: str
    path: str

@dg.asset(
    kinds={"s3"}
)
def import_file_s3(
    config: IngestionFileS3Config,
) -> str:
    s3_path = f"s3://{config.bucket}/{config.path}"
    return s3_path

@dg.asset(
    kinds={"duckdb"}
)
def duckdb_table_s3(
    complexDatabase: DuckDBResource,
    import_file_s3: str
):
    table_name = "raw_s3_data"
    with complexDatabase.get_connection() as conn:
        table_query = f"""
            create table if not exists {table_name} (
                date date,
                share_price float,
                amount float,
                spend float,
                shift float,
                spread float
            ) 
        """
        conn.execute(table_query)
        conn.execute(f"copy {table_name} from '{import_file_s3}'(format_csv, header)")