To create a PostgreSQL container with an initial user and database using Docker compose, you configure the environment variables POSTGRES_USER, POSTGRES_PASSWORD and POSTGRES_DB inside your YAML file

### Import Details to Remember
- First Run Only: PostgreSQL only create the custom user and database and its initialization scripts the very first time the container starts. If you alreay have data in the postgres_data volume, new credentials will be ignored.

If you change, you have to wipe the existing volume

### psycopg2.errors.InsufficientPrivilege: permission denied for schema public postgres
This error frequently occurs in PostgreSQL 15 and newer, which changed default security settings by revoking automatic CREATE privileges from non-owner users on the public schema to prevent security vulnerabilities.

• Make the User the Database Owner: If this user manages the entire application database (e.g., executing Django or SQLAlchemy migrations), transfer ownership of the database to them:
```sql
ALTER DATABASE your_database_name OWNER TO your_username;
```
• Change Schema Ownership: Keep the database owner intact but make the user the direct owner of the public schema:
```sql
ALTER SCHEMA public OWNER TO your_username;
```
• Grant All Privileges: To quickly resolve all permissions (including data manipulation and structure modification) within that schema:
```sql
GRANT ALL PRIVILEGES ON SCHEMA public TO your_username;
```