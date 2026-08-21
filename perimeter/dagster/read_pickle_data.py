import pickle

file_path = "dagster-etl/.tmp_dagster_home_ztk2_8jk/storage/asteroids_partition/2026-08-20"  # e.g., /path/to/dagster_home/storage/asset_name

with open(file_path, "rb") as f:
  data = pickle.load(f)

print(data)
