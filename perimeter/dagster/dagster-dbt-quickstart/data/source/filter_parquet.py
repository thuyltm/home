import pandas as pd

# Load parquet file
df = pd.read_parquet('/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/source/taxi_trips.parquet')

# Ensure the column is datetime type
df['tpep_pickup_datetime'] = pd.to_datetime(df['tpep_pickup_datetime'])

# Define start and end times
start_time = '2026-01-01 00:00:00'
end_time = '2026-01-01 23:59:59'

# Filter using .between()
filtered_df = df[df['tpep_pickup_datetime'].between(start_time, end_time)]

print(filtered_df)
