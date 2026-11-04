#! /bin/sh
gcloud services enable compute.googleapis.com

gcloud compute instances create free-vm-instance \
    --zone=us-central1-a \
    --machine-type=e2-micro \
    --image-family=debian-12 \
    --image-project=debian-cloud \
    --boot-disk-size=30GB \
    --boot-disk-type=pd-standard \
    --network-interface=network-tier=STANDARD \
    --service-account="thuy-le-thi-minh@articulate-run-306102.iam.gserviceaccount.com" \
    --scopes=cloud-platform

# connect to VM instance via ssh
gcloud compute ssh --zone "us-central1-a" "free-vm-instance" --project "articulate-run-306102"

# establish an ssh tunnel that forward the VM's internal Postgres port to a local port
gcloud compute ssh "free-vm-instance" \
    --zone=us-central1-a \
    --project=articulate-run-306102 \
    -- -N -L 5433:localhost:5432
# this command find the PID using the port 5432
sudo lsof -i :5433
# connect using the local machine's psql client
psql -h localhost -p 5433 -U postgres