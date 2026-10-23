import dagster as dg
from dagster_dbt import build_dbt_asset_selection

from dagster_dbt_quickstart.defs.assets.dbt import dbt_analytics, incremental_dbt_models
from dagster_dbt_quickstart.defs.partitions import monthly_partition, weekly_partition, daily_partition

trips_by_week = dg.AssetSelection.assets("trips_by_week")
adhoc_request = dg.AssetSelection.assets("adhoc_request")

# Pass in the dbt_analytics definition and a string that selects stg_trips and all dbt models downstream of it.
dbt_trips_selection = dg.define_asset_job(
    name="dbt_trips_selection",  
    selection=build_dbt_asset_selection([dbt_analytics], "stg_trips+")
)

trip_update_job = dg.define_asset_job(
    name="trip_update_job",
    selection=build_dbt_asset_selection([dbt_analytics], "stg_trips").downstream()
)

# the build_dbt_asset_selection function only selects dbt models and not Dagster assets
# using Dagster's downstream functionality to select all downstream assets
location_metrics_update_job = dg.define_asset_job(
    name="location_metrics_update_job",
    selection=build_dbt_asset_selection([dbt_analytics], "location_metrics").downstream()
)

daily_metrics_update_job = dg.define_asset_job(
    name="daily_metrics_update_job",
    partitions_def=daily_partition,
    selection=build_dbt_asset_selection([incremental_dbt_models], "daily_metrics").downstream()
)

# the AssetSelection.all() call load your dbt project into your Dagster project
# Dagster will recognize and treat all of your dbt models as assets
trip_update_job_exclude = dg.define_asset_job(
    name="trip_update_job_exclude",
    partitions_def=monthly_partition,   
    selection=dg.AssetSelection.all() - adhoc_request - trips_by_week
    - build_dbt_asset_selection([dbt_analytics], "stg_trips+")
    - build_dbt_asset_selection([incremental_dbt_models], "daily_metrics+")
)

weekly_update_job = dg.define_asset_job(
    name="weekly_update_job", partitions_def=weekly_partition, selection=trips_by_week
)

adhoc_request_job = dg.define_asset_job(
    name="adhoc_request_job", selection=adhoc_request
)

