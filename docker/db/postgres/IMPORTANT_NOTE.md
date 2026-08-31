To create a PostgreSQL container with an initial user and database using Docker compose, you configure the environment variables POSTGRES_USER, POSTGRES_PASSWORD and POSTGRES_DB inside your YAML file

### Import Details to Remember
- First Run Only: PostgreSQL only create the custom user and database and its initialization scripts the very first time the container starts. If you alreay have data in the postgres_data volume, new credentials will be ignored.

If you change, you have to wipe the existing volume