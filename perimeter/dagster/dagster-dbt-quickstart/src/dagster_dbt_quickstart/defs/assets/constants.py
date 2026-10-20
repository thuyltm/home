import os

current_path = os.getcwd()

TAXI_ZONES_FILE_TO_WRITE_PATH = os.path.join(current_path, "/data/raw/taxi_zones.csv")
TAXI_TRIPS_FILE_TO_WRITE_PATH = os.path.join(current_path, "data/raw/taxi_trips.parquet")

TRIPS_BY_AIRPORT_FILE_WRITE_PATH = os.path.join(current_path, "data/raw/trips_by_airport.csv")
TRIPS_BY_WEEK_FILE_WRITE_PATH = os.path.join(current_path, "data/raw/trips_by_week.csv")
MANHATTAN_STATS_FILE_WRITE_PATH = os.path.join(current_path, "data/raw/manhattan_stats.geojson")
MANHATTAN_MAP_FILE_WRITE_PATH = os.path.join(current_path, "data/raw/manhattan_map.png")
REQUEST_DESTINATION_PLOT_FILE_WRITE_PATH = os.path.join(current_path, "data/raw/result_request_plot.png")
REQUEST_DESTINATION_RESULT_FILE_WRITE_PATH = os.path.join(current_path, "data/raw/result_request.csv")

DATE_FORMAT = "%Y-%m-%d %H:%M:%S"
START_DATE = "2026-01-01 00:00:00"
END_DATE = "2026-02-01 23:59:59"

AIRPORT_TRIPS_PLOT_FILE_WRITE_PATH = os.path.join(current_path, "data/raw/airport_trips_plot.png")
AIRPORT_TRIPS_RESULT_FILE_WRITE_PATH = os.path.join(current_path, "data/raw/airport_trips.csv")