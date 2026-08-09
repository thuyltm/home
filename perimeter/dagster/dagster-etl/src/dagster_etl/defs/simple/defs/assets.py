

import csv

import dagster as dg

from pathlib import Path
from dagster_duckdb import DuckDBResource

class IngestionFileConfig(dg.Config):
    path: str

@dg.asset()
def import_file(config: IngestionFileConfig) -> str:
    file_path = (
        Path(__file__).absolute().parent / f"../../../data/source/{config.path}"
    )
    return str(file_path.resolve())

@dg.asset(
    kinds={"duckdb"},
)
def duckdb_table(simpleDatabase: DuckDBResource, import_file):
    table_name = "raw_data"
    with simpleDatabase.get_connection() as conn:
        table_query = f"""
            create table if not exists {table_name} (
                date date,
                share_price float,
                amount float,
                spend float,
                shift float,
                spread float
            )"""
        conn.execute(table_query)
        conn.execute(f"copy {table_name} from '{import_file}';")  

@dg.asset_check(
    asset=import_file,
    blocking=True,#which prevents downstream assets from executing if the check fails
    description="Ensure file contains no zero value shares",
)
def not_empty(
    import_file
) -> dg.AssetCheckResult:
    with open(import_file, mode="r", encoding="utf-8") as file:
        reader = csv.DictReader(file)
        data = (row for row in reader)
        for row in data:
            if float(row["share_price"]) <=0:
                return dg.AssetCheckResult(
                    passed=False,
                    metadata={"share' is below O": row}
                )
    return dg.AssetCheckResult(
        passed=True,
    )
