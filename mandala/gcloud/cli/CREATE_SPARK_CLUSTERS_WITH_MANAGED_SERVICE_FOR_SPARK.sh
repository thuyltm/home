#! /bin/sh
###############################################################################
# 1. You can list the active account name
###############################################################################
gcloud auth list
###############################################################################
# 2. You can list the project ID
###############################################################################
gcloud config list project
#[core]
#project=qwiklabs-gcp-04-75c59ef38307
###############################################################################
# 3. Confirm Cloud Dataproc API is enabled
###############################################################################
gcloud services enable dataproc.googleapis.com
gcloud services list --enabled --filter="name:dataproc.googleapis.com"
###############################################################################
# 4. Grant a necessary IAM role to the Compute Engine service account
###############################################################################
gcloud iam service-accounts list
# DISPLAY NAME: Compute Engine default service account
# EMAIL: 400477736790-compute@developer.gserviceaccount.com
gcloud projects add-iam-policy-binding qwiklabs-gcp-01-19fb0e057ee1 \
  --member="serviceAccount:636058934042-compute@developer.gserviceaccount.com" \
  --role="roles/storage.admin"
################################################################################
# Task 1: Create a cluster
################################################################################
gcloud dataproc clusters create example-cluster \
  --region=us-central1 \
  --zone=us-central1-a \
  --master-machine-type=e2-standard-2 \
  --master-boot-disk-type=pd-standard \
  --master-boot-disk-size=30GB \
  --worker-machine-type=e2-standard-2 \
  --worker-boot-disk-type=pd-standard \
  --worker-boot-disk-size=30GB \
  --num-workers=2
###############################################################################
# Taks 2: Submit a Job
###############################################################################
gcloud dataproc jobs submit spark \
    --cluster=example-cluster \
    --region=us-central1 \
    --class=org.apache.spark.examples.SparkPi \
    --jars=file:///usr/lib/spark/examples/jars/spark-examples.jar \
    -- 1000
##############################################################################
# Task 3: Update a cluster configuration
##############################################################################
gcloud dataproc clusters update example-cluster \
  --region=us-central1 \
  --num-workers=4