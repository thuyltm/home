import os

import dlt
import requests

@dlt.source
def nasa_neo_source(start_date: str, end_date: str, api_key: str):
    @dlt.resource
    def fetch_neo_data():
        url = "https://api.nasa.gov/neo/rest/v1/feed"
        params = {
            "start_date": start_date,
            "end_date": end_date,
            "api_key": api_key,
        }
        response = requests.get(url, params=params)
        response.raise_for_status()

        data = response.json()

        for neo in data["near_earth_objects"][start_date]:
            neo_data = {
                "id": neo["id"],
                "name": neo["name"],
                "absolute_magnitude_h": neo["absolute_magnitude_h"],
                "is_potentially_hazardous": neo["is_potentially_hazardous_asteroid"],
            }
            yield neo_data
    return fetch_neo_data

pipeline=dlt.pipeline(
    pipeline_name="nasa_pipeline",
    destination=dlt.destinations.duckdb("data.duckdb"),
    dataset_name="nasa_neo",
)

pipeline.run(
    nasa_neo_source(
        start_date="2026-08-24",
        end_date="2026-08-25",
        api_key="KzY7FJKve4kv3v5h6MaFfjsychngZJCZfU1AN7Br",
    )
)
    