[Guide](https://docs.slingdata.io/)
Many people replicate data from an OLTP database (like Postgres or MySQL) into a data warehouse (such as Snowflake or Redshift)

Because schema drift and type mismatches cause performance bottlenecks and data consistency issues, moving data between systems requires careful handling and thoughtful architecture.


| Database Type | Query         | Examples                          |
| --------------| --------------| --------------------------------- |
| Relational    | SQL           | Postgres, MySQL, Oracle, SQLite   |
| NoSQL         | Varies        | MongoDB, Redis, DynamoDB          |
| Graph         | Neo4j, Gremlin| Neo4j, Neptune, ArangoDB          |
| Vector        | Semantic      | FAISS, Pinecone, Weaviate         |
| Time-series   | SQL-like      | InfluxDB, Premetheus              |

#### How many database replication methods are there?
##### Full refresh replication
IF we rely solely on full refreshes, we have to run the entire extraction process every time. THIS can be prohibitively expensive for large tables AND can strain the source database, impacting performance during the sync
##### Incremental replication
You can optimize full refreshes by filtering the data. Usually THIS involves querying based on time columns or incrementing ids. The checkpoint of the last query is THEN maintained by the ETL service SO it knows where begin replication
##### Change data capture (CDC)
Relational databases like Postgres maintain a log of changes, this is the Write-Ahead Log (WAL), which records inserts, updates, and deletes for recovery and replication purposes.
1. CDC logs are retained only for a few days
2. CDC doesn't provide full historical context,  so it can't be used alone to initialize a replica
#### Building database replication systems
A tool may perform a full refresh to establish the initial snapshot of the table and then switch to CDC to capture all changes moving forward