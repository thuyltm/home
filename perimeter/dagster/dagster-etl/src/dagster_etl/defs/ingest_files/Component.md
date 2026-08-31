So far, we've seen how Dagster can work with frameworks like dlt and Sling to build robust pipeline. However, __coordinating across frameworks can be difficult__. In order to simplify this development process, __Dagster unifies this all within components__. This goal is to reduce the amount of code and developing these type of integration much easy

Dagster Components provide __a class-based interface (Component) that encapsulating__ complex logic into resuable, declarative building blocks.

For example: a component can manage our dlt connection between Postgres and DuckDB. This eliminates the need for low-level setup and lets you focus solely on the relevant inputs and outputs for your use case. This indicates that users can __provide a few configuration values, and let the component handle the orchestration__.

1. To get a list of components in your dagster project
```sh
dg list defs
```
2. With the aim of viewing the available components in your environment
```sh
dg list components
```