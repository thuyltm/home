1. dbt seed: loads CSV files located in the seed-paths directory of your project into your data warehouse
2. dbt compile: generates executable SQL from your project source files. it requires a live data warehouse connection to fetch metadata
3. dbt run: compiles and runs your project
4. dbt test: compiles and tests your project
5. dbt build: compiles, runs, and tests your project
6. dbt docs generate: generates your project's documentation
7. dbt docs serve: starts a webserver on port 8080 to serve your documentation locally and opens the documentation site in your default browser
8. dbt parse: reads and validates your project files, constructs the Directed Acyclic Graph (DAG) and generates a manifest.json artifact without connecting to your data warehouse or compile final SQL