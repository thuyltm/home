import dagster as dg

from dagster_etl.defs.complex.defs.sling.jobs import (
    orders_refresh_job,
    postgres_refresh_job,
)

postgres_refresh_schedule = dg.ScheduleDefinition(
    job=postgres_refresh_job,
    cron_schedule="* * * * *", # every minute
)

orders_refresh_schedule = dg.ScheduleDefinition(
    job=orders_refresh_job,
    cron_schedule="*/2 * * * *", # minite hour dayofmonth month dayofweek
)