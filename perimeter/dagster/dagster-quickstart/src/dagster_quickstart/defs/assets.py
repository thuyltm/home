import pandas as pd
from dagster_duckdb import DuckDBResource

import dagster as dg

sample_data_file = "src/dagster_quickstart/defs/data/sample_data.csv"
processed_data_file = "src/dagster_quickstart/defs/data/processed_data.csv"

@dg.asset
def processed_data():
    ## Read data from the CSV
    df = pd.read_csv(sample_data_file)
    ## Add an age_group column based on the value of age
    df["age_group"] = pd.cut(
        df["age"], bins=[0, 30, 40, 100], labels=["Young", "Middle", "Senior"]
    )
    ## Save processed data
    df.to_csv(processed_data_file, index=False)
    return "Data loaded successfully"

@dg.asset
def customers(duckdb: DuckDBResource):
    url = "src/dagster_quickstart/defs/data/raw_customers.csv"
    table_name = "customers"
    with duckdb.get_connection() as conn:
        conn.execute(
            f"""
            create or replace table {table_name} as (
                select * from read_csv_auto('{url}')
            )
            """
        )

@dg.asset
def orders(duckdb: DuckDBResource):
    url = "src/dagster_quickstart/defs/data/raw_orders.csv"
    table_name = "orders"
    with duckdb.get_connection() as conn:
        conn.execute(
            f"""
            create or replace table {table_name} as (
                select * from read_csv_auto('{url}')
            )
            """
        )

@dg.asset
def payments(duckdb: DuckDBResource):
    url = "src/dagster_quickstart/defs/data/raw_payments.csv"
    table_name = "payments"
    with duckdb.get_connection() as conn:
        conn.execute(
            f"""
            create or replace table {table_name} as (
                select * from read_csv_auto('{url}')
            )
            """
        )

@dg.asset(
    deps=["customers", "orders", "payments"],
)
def orders_aggregation(duckdb: DuckDBResource):
    table_name = "orders_aggregation"
    with duckdb.get_connection() as conn:
        conn.execute(
            f"""
            create or replace table {table_name} as (
                select  c.id as customer_id,
                        c.first_name,
                        c.last_name,
                        count(distinct o.id) as total_orders,
                        count(distinct p.id) as total_payments,
                        coalesce(sum(p.amount), 0) as totatl_amount_spent
                from customers c
                left join orders o on c.id = o.user_id
                left join payments p on o.id = p.order_id
                group by 1, 2, 3
            );
            """
        )

@dg.asset_check(asset="orders_aggregation")
def orders_aggregation_check(duckdb: DuckDBResource) -> dg.AssetCheckResult:
    table_name = "orders_aggregation"
    with duckdb.get_connection() as conn:
        row_count = conn.execute(f"select count(*) from {table_name}").fetchone()[0]
    if row_count == 0:
        return dg.AssetCheckResult(
            passed=False, metadata={"message": "Order aggregation check failed."}
        )
    return dg.AssetCheckResult(
        passed=True, metadata={"message": "Order aggregation check passed."}
    )

