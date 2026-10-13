import dagster as dg
from dagster_dbt import build_dbt_asset_selection

from dagster_dbt_quickstart.defs.assets.dbt import dbt_analytics
from dagster_dbt_quickstart.defs.partitions import monthly_partition, weekly_partition

trips_by_week = dg.AssetSelection.assets("trips_by_week")
adhoc_request = dg.AssetSelection.assets("adhoc_request")

trip_update_job = dg.define_asset_job(
    name="trip_update_job",
    partitions_def=monthly_partition,
    selection=build_dbt_asset_selection([dbt_analytics], "stg_trips+")
)

zone_update_job = dg.define_asset_job(
    name="zone_update_job",
    partitions_def=monthly_partition,
    selection=build_dbt_asset_selection([dbt_analytics], "stg_zone+")
)

location_metrics_update_job = dg.define_asset_job(
    name="location_metrics_update_job",
    selection=build_dbt_asset_selection([dbt_analytics], "location_metrics+")
)

daily_metrics_update_job = dg.define_asset_job(
    name="daily_metrics_update_job",
    selection=build_dbt_asset_selection([dbt_analytics], "daily_metrics+")
)

trip_zone_update_job = dg.define_asset_job(
    name="trip_zone_update_job",
    selection=dg.AssetSelection.all()
    - build_dbt_asset_selection([dbt_analytics], "stg_trips+")
    - build_dbt_asset_selection([dbt_analytics], "stg_zone+")
)

weekly_update_job = dg.define_asset_job(
    name="weekly_update_job", partitions_def=weekly_partition, selection=trips_by_week
)

adhoc_request_job = dg.define_asset_job(
    name="adhoc_request_job", selection=adhoc_request
)

