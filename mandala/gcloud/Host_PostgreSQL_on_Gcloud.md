To host PostgreSQL on Google Cloud as cheaply as possible, you should avoid the fully managed Cloud SQL and instead run PostgreSQL inside a self-managed Compute Engine VM

#### Create the Free VM Instance
1. Navigate to __Compute Engine > VM instances__ and click Create Instance
2. Configure the following specific settings to stay free
- Region: us-central1, us-east1, us-west1 (the free tier is strictly limited to these regions)
- Machine Series: E2
- Machine Type: e2-micro (2 vCPUs, 1 GB RAM)
- Boot Disk: Select Ubuntu or Debian. Change the disk type to __Standard Persistent Disk__ and set the size to 30 GB or less (Do not select SSD if you want it to remain 100% free)

#### Open the PostgreSQL 
By default, Google Cloud blocks all incoming traffic to your VM. You need to open port 5432 so you can connect to your database
1. Search for VPC Network in the top search bar and go to Firewall
2. Click Create Firewall Rule
3. Set the following parameters:
- Name: allow-postgres
- Targets: All instances in the network
- Source IPv4 range: _0.0.0.0/0_
- Protocols and ports: Check Specified protocols and ports, check TCP, and type 5432

#### Install and Configure PostgreSQL
Once the VM is running, click the SSH button in the Compute Engine console to open a terminal
```sh
sudo apt update && sudo apt upgrade -y
sudo apt install postgresql postgresql-contrib -y
```
1. Set a password for the default database user
```sh
sudo -i -u postgres
psql
ALTER USER postgres WITH PASSWORD 'your_secure_password_for_user_postgres';
\q
exit
```
__OR__
```sh
sudo -i -u postgres psql -c "ALTER USER postgres PASSWORD '123456';"
```
2. Configure Postgres to accept external connections
By default, Postgres only listens to local traffic. You must tell it to listen to external connections:

- Open the main configuration file
```sh
sudo nano /etc/postgreslql/16/main/postgresql.conf
```
Change __listen_addresses='*'__
- Authoriza users to authenticate remotely
```sh
sudo nano /etc/postgresql/16/main/pg_hba.conf
```
Add the following line to allow connections from any IP addressusing password encryption
```sh
host all all 0.0.0.0/0 scram-sha-256
```
- Restart Postgres to apply changes
```sh
sudo systemctl restart postgresql
```

#### Connect to your database

- Host: Your_VM_External_IP
- Port: 5432
- User: postgres
- Password: 123456
```sh
psql postgresql://test_user:test_pass@35.209.239.103:5432/test_db
```
#### ssh into the VM instance
```sh
gcloud compute ssh --zone "us-central1-a" "free-vm-instance" --project "articulate-run-306102"
```
#### Establish an ssh tunnel that forward the VM's internal Postgres port to a local port
```sh
gcloud compute ssh "free-vm-instance" \
    --zone=us-central1-a \
    --project=articulate-run-306102 \
    -- -N -L 5433:localhost:5432
```
#### Your initialization SQL commands go here
```sh
# 1. Create the database
CREATE DATABASE test_db;
# 2. Create the user test_user
\du
CREATE ROLE test_user WITH LOGIN PASSWORD 'test_pass';
# 3. Grant privileges to the user for that database
GRANT ALL PRIVILEGES ON DATABASE test_db TO test_user;
# 4. Change the owner of database 
ALTER DATABASE test_db OWNER TO test_user;
\l
\dn+ public
```
#### Logging
```sh
tail -f /var/log/postgresql/postgresql-*.log
```
