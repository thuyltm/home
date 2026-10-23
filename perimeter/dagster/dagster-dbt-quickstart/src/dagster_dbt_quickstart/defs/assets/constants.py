from pathlib import Path

# 1. Get the directory containing the current script
# .resolve() makes the path absolute and resolves any symlinks
parent_dir = Path('.').joinpath("data").resolve()

TAXI_ZONES_FILE_TO_WRITE_PATH = parent_dir.joinpath("raw/taxi_zones.csv").resolve()
TAXI_TRIPS_FILE_TO_WRITE_PATH = parent_dir.joinpath("raw/taxi_trips.parquet").resolve()

TRIPS_BY_AIRPORT_FILE_WRITE_PATH = parent_dir.joinpath("raw/trips_by_airport.csv").resolve()
TRIPS_BY_WEEK_FILE_WRITE_PATH = parent_dir.joinpath("raw/trips_by_week.csv").resolve()
MANHATTAN_STATS_FILE_WRITE_PATH = parent_dir.joinpath("raw/manhattan_stats.geojson").resolve()
MANHATTAN_MAP_FILE_WRITE_PATH = parent_dir.joinpath("raw/manhattan_map.png").resolve()
REQUEST_DESTINATION_PLOT_FILE_WRITE_PATH = parent_dir.joinpath("raw/result_request_plot.png").resolve()
REQUEST_DESTINATION_RESULT_FILE_WRITE_PATH = parent_dir.joinpath("raw/result_request.csv").resolve()

DATE_FORMAT = "%Y-%m-%d %H:%M:%S"
START_DATE = "2026-01-01 00:00:00"
END_DATE = "2026-02-01 23:59:59"

AIRPORT_TRIPS_PLOT_FILE_WRITE_PATH = parent_dir.joinpath("raw/airport_trips_plot.png").resolve()
AIRPORT_TRIPS_RESULT_FILE_WRITE_PATH = parent_dir.joinpath("raw/airport_trips.csv").resolve()