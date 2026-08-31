import dagster as dg

postgres_refresh_job = dg.define_asset_job(
    "postgres_refresh",
    selection=[
        dg.AssetKey(["target", "test_schema","customers"]),
        dg.AssetKey(["target", "test_schema", "products"]),
        dg.AssetKey(["target", "test_schema", "orders"]),
    ],
)

orders_refresh_job = dg.define_asset_job(
    "orders_refresh",
    selection=[
        dg.AssetKey(["target", "test_schema", "orders"]),
    ],
)