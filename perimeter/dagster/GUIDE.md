https://docs.dagster.io/
[Guide](https://courses.dagster.io/courses/take/dagster-etl/multimedia/65303173-about-this-course)
[Github](https://github.com/dagster-io/project-dagster-university/tree/main/dagster_university)

Dagster's programming model is built around Assets, which are organized with __Definitions__
- __Definitions__ act as containers for all Dagster entities in your project
- __Assets__ are the core building blocks of Dagster. They represent the entities in your data platform, such as database tables, machine learning models or AI processes, while the dependencies between them define project's lineage.

    An asset represents a logical unit of data such as table, dataset, or machine learning model. Assets can have dependencies on other assets, forming the data lineage for your pipeline.
- __Resources__ Reources are Dagster objects much like assets, but they are not executed. Typically, resouces are reusable objects that supply external context, such as database connections, API clients, or configuration settings. A single resource can be shared accross many different Dagster objects

    For example: You load that data into DuckDB, an analytical database, so your assets will be representations of DuckDB tables. Since the same database will be used across the system, you can use a resource to centralize the connection in a single object that can be shared across multiple Dagster objects
- __Automation__ Dagster supports both scheduled and event-driven pipelines. 
- __Custom component__ With custom components, you can define your own specific use cases
- __Partition__ Partitions allow you to break down your data into smaller, logical segments. Configured an asset with partitions allows you to:
    - Materialize and update only specific partitions, avoiding unneccessary reprocessing of unchanged data
    - Recover from failures without rerunning the entire pipeline
- __Dynamic Partition__ what about cases where the upstream data doesn't follow a fixed pattern? For example, what if the data is grouped by something like customer name? In this case, the set of partitions may evolve over time as new customers are added

1. Scaffold a new Dagster project
```sh
uvx create-dagster@latest project dagster-tutorial
```
2. Scaffold an assets file with the _dg scaffold_ command
```sh
dg scaffold defs dagster.asset assets.py
```
In dagster, all assets need to be associated with a top-level Definitions object in order to be deployed
```python
@definitions
def defs():
    return load_from_defs_folder(project_root=Path(__file__).parent.parent.parent)
```
3. Start the Dagster webserver
```sh
dg dev
```
4. Scaffold a resources file
```sh
dg scaffold defs dagster.resources resources.py
```
5. Confirm the assets and resources are configured correctly
```sh
dg check defs
```
6. To run the pipeline, click the Assets tab/ Materialize
7. Scaffold a new schedule object
```sh
dg scaffold defs dagster.schedule schedules.py
```
8. Scaffold a custom component
```sh
dg scaffold component Tutorial
```
9. List your components
```sh
dg list components
```
10. Scaffold defintions from component:
```sh
 dg scaffold defs dagster_quickstart.components.tutorial.Tutorial tutorial
```