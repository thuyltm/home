#! /bin/sh
# GUIDE: https://docs.cloud.google.com/iam/docs/granting-changing-revoking-access
gcloud config list
# project = articulate-run-306102
# gcloud RESOURCE_TYPE get-iam-policy RESOURCE_ID --format=FORMAT > PATH
gcloud projects get-iam-policy articulate-run-306102 --format=json > policy.json
# gcloud RESOURCE_TYPE add-iam-policy-binding RESOURCE_ID \
#    --member=PRINCIPAL --role=ROLE_NAME \
#    --condition=CONDITION
gcloud projects add-iam-policy-binding articulate-run-306102 \
    --member="serviceAccount:thuy-le-thi-minh@articulate-run-306102.iam.gserviceaccount.com" \
    --role="roles/iam.workloadIdentityPoolAdmin"
gcloud projects add-iam-policy-binding articulate-run-306102 \
    --member="serviceAccount:thuy-le-thi-minh@articulate-run-306102.iam.gserviceaccount.com" \
    --role="roles/iam.serviceAccountAdmin"
#gcloud RESOURCE_TYPE remove-iam-policy-binding RESOURCE_ID
#   --member=PRINCIPAL --role=ROLE_NAME