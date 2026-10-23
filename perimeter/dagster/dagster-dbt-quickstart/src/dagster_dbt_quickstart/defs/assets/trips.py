from io import BytesIO
from pathlib import Path

import dagster as dg
import pandas as pd
from dagster_duckdb import DuckDBResource
from . import constants
import os

@dg.asset
def list_and_process_files(context: dg.AssetExecutionContext):
    # Get filenames of all files in the current directory
    files = [f.name for f in Path('.').iterdir() if f.is_file()]
    context.log.info(files)
    folders = [f.name for f in Path('.').iterdir() if f.is_dir()]
    context.log.info(folders)
    data_folders = Path('.').joinpath("data").resolve()
    sub_data_folders = [f.name for f in data_folders.iterdir() if f.is_dir()]
    context.log.info(sub_data_folders)
    for file_name in sub_data_folders:
        folder_path = Path('.').joinpath("data", file_name).resolve()
        for item in folder_path.iterdir():
            if item.is_file():
                context.log.info(item.name)

@dg.asset(
    group_name="raw_files",
    kinds={"python"},
)
def taxi_zones_file() -> dg.MaterializeResult:
    parent_dir = Path('.').joinpath("data").resolve()
    file_to_read = parent_dir.joinpath("source/taxi_zones.csv").resolve()
    # 1. Open the source file in read mode ('r')
    with open(file_to_read, 'rb') as source:
        # source.read() outputs a 'str' containing the file's content
        file_content = source.read() 
    with open(
        constants.TAXI_ZONES_FILE_TO_WRITE_PATH, "wb"
    ) as output_file:
        output_file.write(file_content)
    num_rows = len(pd.read_csv(BytesIO(file_content)))
    return dg.MaterializeResult(
        metadata={"Number of records": dg.MetadataValue.int(num_rows)}
    )

@dg.asset(
    deps=["taxi_zones_file"],
    group_name="ingested",
    kinds={"duckdb"},
)
def taxi_zones(context: dg.AssetExecutionContext, database: DuckDBResource):
    query = f"""
        create or replace table zones as (
            select
                OBJECTID,
                zone,
                borough,
                the_geom,
                LOCATIONID,
            from '{constants.TAXI_ZONES_FILE_TO_WRITE_PATH}'
        );
    """
    with database.get_connection() as conn:
        conn.execute(query)

@dg.asset(
    group_name="raw_files",
    kinds={"duckdb"},
)
def taxi_trips_file(context: dg.AssetExecutionContext) -> dg.MaterializeResult:
    parent_dir = Path('.').joinpath("data").resolve()
    file_to_read = parent_dir.joinpath("source/taxi_trips.parquet").resolve()

    # 1. Open the source file in read mode ('r')
    with open(file_to_read, 'rb') as source:
        # source.read() outputs a 'str' containing the file's content
        file_content = source.read() 

    with open(
        constants.TAXI_TRIPS_FILE_TO_WRITE_PATH, "wb"
    ) as output_file:
        output_file.write(file_content)

    num_rows = len(pd.read_parquet(BytesIO(file_content)))
    return dg.MaterializeResult(
        metadata={"Number of records": dg.MetadataValue.int(num_rows)}
    )

@dg.asset(
    deps=["taxi_trips_file"],
    group_name="ingested",
    kinds={"duckdb"},
)
def taxi_trips(context: dg.AssetExecutionContext, database: DuckDBResource):
    query = f"""
        create table if not exists trips (
            VendorID integer, tpep_pickup_datetime timestamp, tpep_dropoff_datetime timestamp,
            passenger_count bigint, trip_distance double, RatecodeID bigint,
            store_and_fwd_flag varchar, PULocationID integer, DOLocationID integer,
            payment_type bigint, fare_amount double, extra double, mta_tax double,
            tip_amount double, tolls_amount double, improvement_surcharge double,
            total_amount double, congestion_surcharge double, Airport_fee double,
            cbd_congestion_fee double
        );

        insert into trips
        select
            VendorID, tpep_pickup_datetime, tpep_dropoff_datetime,
            passenger_count, trip_distance, RatecodeID,
            store_and_fwd_flag, PULocationID, DOLocationID,
            payment_type, fare_amount, extra, mta_tax,
            tip_amount, tolls_amount, improvement_surcharge,
            total_amount, congestion_surcharge, Airport_fee,
            cbd_congestion_fee
        from '{constants.TAXI_TRIPS_FILE_TO_WRITE_PATH}';
    """

    with database.get_connection() as conn:
        conn.execute(query)