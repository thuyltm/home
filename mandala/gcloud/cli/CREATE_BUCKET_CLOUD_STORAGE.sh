#! /bin/sh
gcloud storage buckets create gs://my-city-vt-bucket \
    --location=us-central1 \
    --default-storage-class=standard \
    --uniform-bucket-level-access
######################################################
# Delete an empty bucket
######################################################
gcloud storage buckets delete gs://BUCKET_NAME
######################################################
# Delete a non-empty bucket
######################################################
gcloud storage buckets delete gs://my-city-vt-bucket
###############################################################################################
# Bucket is not empty. To delete all objects and then delete bucket, use: gcloud storage rm -r
###############################################################################################
gcloud storage rm -r gs://my-city-vt-bucket