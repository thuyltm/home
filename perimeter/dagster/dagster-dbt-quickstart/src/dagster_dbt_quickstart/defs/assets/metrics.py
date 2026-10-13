import base64

import dagster as dg
import geopandas as gpd
import matplotlib.pyplot as plt
import pandas as pd
from dagster_duckdb import DuckDBResource
from dagster_dbt_quickstart.defs.assets import constants
from dagster_dbt_quickstart.defs.partitions import weekly_partition

@dg.asset(
    deps=[dg.AssetKey(["taxi_trips"])],
    partitions_def=weekly_partition,
    kinds={"duckdb"},
)
def trips_by_week(context: dg.AssetExecutionContext, database: DuckDBResource):
    period_to_fetch = context.partition_key
    query = f"""
        select VendorID, total_amount, trip_distance, passenger_count
        from trips
    """
    with database.get_connection() as conn:
        data_for_week = conn.execute(query).fetch_df()
    aggregate = (
        data_for_week.agg(
            {
                "VendorID": "count",
                "total_amount": "sum",
                "trip_distance": "sum",
                "passenger_count": "sum",
            }
        )
        .rename({"VendorID": "num_trips"})
        .to_frame()
        .T
    )
    # clean up the formatting of the dataframe
    aggregate["period"] = period_to_fetch
    aggregate["num_trips"] = aggregate["num_trips"].astype(int)
    aggregate["passenger_count"] = aggregate["passenger_count"].astype(int)
    aggregate["total_amount"] = aggregate["total_amount"].round(2).astype(float)
    aggregate["trip_distance"] = aggregate["trip_distance"].round(2).astype(float)
    aggregate = aggregate[
        ["period", "num_trips", "total_amount", "trip_distance", "passenger_count"]
    ]

    try:
        # If the file already exists, append to it, but replace the existing month's data
        existing = pd.read_csv(constants.TRIPS_BY_WEEK_FILE_WRITE_PATH)
        existing = existing[existing["period"] != period_to_fetch]
        existing = pd.concat([existing, aggregate]).sort_values(by="period")
        existing.to_csv(constants.TRIPS_BY_WEEK_FILE_WRITE_PATH, index=False)
    except FileNotFoundError:
        aggregate.to_csv(constants.TRIPS_BY_WEEK_FILE_WRITE_PATH, index=False)

@dg.asset(
    deps=[dg.AssetKey(["taxi_trips"]), dg.AssetKey(["taxi_zones"])],
    key_prefix="manhattan",
    kinds={"duckdb"},
)
def manhattan_stats(database: DuckDBResource):
    query = """
        select
            zones.zone,
            zones.borough,
            zones.the_geom as geometry,
            count(1) as num_trips,
        from trips
        left join zones on trips.PULocationID = zones.LocationID
        where the_geom is not null
        group by zone, borough, the_geom
    """

    with database.get_connection() as conn:
        trips_by_zone = conn.execute(query).fetch_df()

    trips_by_zone["geometry"] = gpd.GeoSeries.from_wkt(trips_by_zone["geometry"])
    trips_by_zone = gpd.GeoDataFrame(trips_by_zone)

    trips_by_zone.to_file(constants.MANHATTAN_STATS_FILE_WRITE_PATH, driver="GeoJSON")
        

@dg.asset(
    deps=[dg.AssetKey(["manhattan", "manhattan_stats"])],
    kinds={"python"},
)
def manhattan_map() -> dg.MaterializeResult:
    """A map of the number of trips per taxi zone in Manhattan."""
    trips_by_zone = gpd.read_file(constants.MANHATTAN_STATS_FILE_WRITE_PATH)

    fig, ax = plt.subplots(figsize=(10, 10))
    trips_by_zone.plot(
        column="num_trips", cmap="plasma", legend=True, ax=ax, edgecolor="black"
    )
    ax.set_title("Number of Trips per Taxi Zone in Manhattan")

    ax.set_xlim(-74.05, -73.90)  # Adjust longitude range
    ax.set_ylim(40.70, 40.82)  # Adjust latitude range

    # Save the image
    plt.savefig(constants.MANHATTAN_MAP_FILE_WRITE_PATH, format="png", bbox_inches="tight")
    plt.close(fig)

    with open(constants.MANHATTAN_MAP_FILE_WRITE_PATH, "rb") as file:
        image_data = file.read()

    # Convert the image data to base64
    base64_data = base64.b64encode(image_data).decode("utf-8")
    md_content = f"![Image](data:image/jpeg;base64,{base64_data})"

    return dg.MaterializeResult(metadata={"preview": dg.MetadataValue.md(md_content)})

@dg.asset(
    deps=["location_metrics"],
)
def airport_trips(database: DuckDBResource) -> dg.MaterializeResult:
    query = """
        select
            zone,
            destination_borough,
            trips
        from location_metrics
        where from_airport
    """
    with database.get_connection() as conn:
        airport_trips = conn.execute(query).fetch_df()

    # Write to CSV file
    airport_trips.to_csv(constants.AIRPORT_TRIPS_RESULT_FILE_WRITE_PATH, index=False)
    # Plot bars
    fig, ax = plt.subplots(figsize=(10, 6))

    # Group data by destination_borough and plot the bar chart
    airport_trips.groupby(["zone", "destination_borough"]).sum()[
        "trips"
    ].unstack().plot(kind="bar", stacked=True, ax=ax)

    # Customize the plot
    ax.set_xlabel("Zone")
    ax.set_ylabel("Number of Trips")
    ax.set_title("Trips from Airport by Destination Borough")
    ax.legend(title="Destination Borough")

    # Save the image
    plt.savefig(constants.AIRPORT_TRIPS_PLOT_FILE_WRITE_PATH, format="png", bbox_inches="tight")
    plt.close(fig)

    # Convert the image data to base64
    with open(constants.AIRPORT_TRIPS_PLOT_FILE_WRITE_PATH, "rb") as file:
        image_data = file.read()

    # Convert the image data to base64
    base64_data = base64.b64encode(image_data).decode("utf-8")
    md_content = f"![Image](data:image/jpeg;base64,{base64_data})"

    return dg.MaterializeResult(metadata={"preview": dg.MetadataValue.md(md_content)})