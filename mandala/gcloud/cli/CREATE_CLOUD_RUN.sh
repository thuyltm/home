#! /bin/sh
gcloud config set run/region us-central1
gcloud config set run/platform managed

######################################################
# Initialize your environment
#####################################################
gclod auth login
gcloud config set project YOUR_PROJECT_ID
#####################################################################
# Enable the Cloud Run Admin API and Cloud Build API
####################################################################
gcloud services enable run.googleapis.com cloudbuild.googleapis.com
######################################################################################
# Deploy a pre-build "Hello World" container
#######################################################################################
gcloud run deploy hello-service \
    --image=us-docker.pkg.dev/cloudrun/container/hello:latest \
    --service-account=thuy-le-thi-minh@articulate-run-306102.iam.gserviceaccount.com \
    --region=us-central1 \
    --min-instances 0 \
    --max-instances 1 \
    --allow-unauthenticated
#############################################################
# Delete
#############################################################
gcloud run services delete hello-service --region us-central1