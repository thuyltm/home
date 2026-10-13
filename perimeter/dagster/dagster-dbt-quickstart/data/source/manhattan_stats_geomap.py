import geopandas as gpd
import matplotlib.pyplot as plt

# 1. Load the GeoJSON file into a GeoDataFrame
gdf = gpd.read_file('/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/raw/manhattan_stats.geojson')

# 2. Set up the matplotlib figure and axis
fig, ax = plt.subplots(figsize=(10, 10))

# 3. Plot directly onto the axis
# gdf.plot(ax=ax, color='lightblue', edgecolor='black', alpha=0.7)
# gdf.plot(ax=ax, edgecolor="black", alpha=0.5, cmap="viridis")
# gdf.plot( ax=ax, edgecolor="black", cmap="plasma", legend=True)
gdf.plot( ax=ax, edgecolor="black", cmap="plasma", legend=True, column="num_trips")

# 4. Customize the plot using standard pyplot commands
ax.set_title("GeoJSON Visualization via GeoPandas", fontsize=14)
ax.set_xlabel("Longitude")
ax.set_ylabel("Latitude")

# Show the plot
plt.savefig("/home/thuy/Documents/Learn/home/perimeter/dagster/dagster-dbt-quickstart/data/raw/manhattan_stats_cmap_legend_column.png")