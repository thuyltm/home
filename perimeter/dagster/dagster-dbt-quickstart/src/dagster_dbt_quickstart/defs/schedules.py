import dagster as dg

from dagster_dbt_quickstart.defs.jobs import trip_update_job, weekly_update_job

trip_update_schedule = dg.ScheduleDefinition(
    job=trip_update_job,
    cron_schedule="*/1 * * * *", # minute hour dayofmonth month dayofweek
)

weekly_update_schedule = dg.ScheduleDefinition(
    job=weekly_update_job,
    cron_schedule="*/1 * * * *",
)