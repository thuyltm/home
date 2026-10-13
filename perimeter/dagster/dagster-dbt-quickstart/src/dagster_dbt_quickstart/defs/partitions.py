import dagster as dg
from datetime import datetime

from dagster_dbt_quickstart.defs.assets import constants

start_date = datetime.strptime(constants.START_DATE, constants.DATE_FORMAT)
end_date = datetime.strptime(constants.END_DATE, constants.DATE_FORMAT)

monthly_partition = dg.MonthlyPartitionsDefinition(
    start_date=start_date, end_date=end_date
)

weekly_partition = dg.WeeklyPartitionsDefinition(
    start_date=start_date, end_date=end_date
)

daily_partition = dg.DailyPartitionsDefinition(start_date, end_date)