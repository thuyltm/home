import dagster as dg
import dagster_etl.defs.complexpartition.defs.jobs as jobs

asset_partitioned_schedule = dg.build_schedule_from_partitioned_job(
    jobs.import_partition_job,
)