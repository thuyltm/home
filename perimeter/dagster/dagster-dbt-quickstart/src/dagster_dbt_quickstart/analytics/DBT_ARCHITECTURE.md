#### What is dbt project structure
1. 
- dbt_project.yml: The main configuration file that defines your project name, version, and folder paths
- setup.py: this python file is used to specify project dependencies for a new python project
- profiles.yml: this file contains connection details for other parts of a data platform, such as a DuckDB or Snowflake database
- pyproject.toml: A file that specifies package core metadata in a static, tool-agnostic way. This file includes a tool.dagster section which references the Python module with your Dagster defintions defined and discoverable at the top level
2. models: the core directory where your SQL transformation files live
3. macros: store reusable SQL functions
4. seeds: holds static CSV reference files loaded directly into your data warehouse
5. tests: contains custom data validation rules and assertion queries
6. snapshots: tracks historical data changes over time