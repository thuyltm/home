import duckdb

import dlt


@dlt.source
def simple_source():
    @dlt.resource
    def load_dict():
        data = [
            {"id": 1, "name": "Alice"},
            {"id": 2, "name": "Bob"},
        ]
        yield data
    return load_dict


pipeline = dlt.pipeline(
    pipeline_name="simple_pipeline",
    destination=dlt.destinations.duckdb("data.duckdb"),
    dataset_name="mydata",
)

pipeline.run(simple_source())