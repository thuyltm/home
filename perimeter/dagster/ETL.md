ETL stands for __Extract, Transform, Load__ and is the process of consolidating data from various upstream sources __into a single destination storage__. To fully leverage this data, it's typically best to bring everything into one centralized location, traditionaly a data warehouse or data lake


| Stage | Types
| --- | --- 
| Extract| Web scraping, exteranl files, database replication, API ingestion, log parsing, message queues 
| Transformation | Data cleaning, normalization, filtering, aggregation, enrichment, joins, deduplication, business logic applications, format conversion
| Load           | Data ware house ingestion (e.g., BigQuery, Snowflake, Redshift), database inserts/updates, data lake storage (e.g., S3, Delta Lake)


Dagster breaks down data pipelines to their individual assets. This provides full lineage and visibility between different assets
