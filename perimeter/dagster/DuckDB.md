DuckDB is an opensource, embedded, relational database management system designed specifically for Online Analytical Processing (OLAP). It operates entirely in-process but is optimized for heavy data crunching rather than transaction processing

**Key Features**
- Columnar Engine: DuckDB stores data column-by-column. This allows it to read only the specific columns needed for a query, drastically speeding up aggregations and filters on large datasets
- Vectorized Query Execution: Data is processed in parallel batches (vectors) instead of row-by-row

**Common Use Cases**
- Local Data Analysis: Querying and aggregating gigabytes of data locally on your laptop without spinning up a massive cloud data warehouse
- Data Pipeline Processing: Serving as a fast, lightweight transformation engine
- Interactive BI Applications: Powering fast dashboards that need to compute complex statics on the fly without database latency
- Use Familiar SQL logic
### Install
```sh
curl https://install.duckdb.org | sh
```