#! /bin/sh
################################################################################################
# Enable the IAM, Resource Manager, Service Account Credentials, and Security Token Service APIs
################################################################################################
gcloud services enable iam.googleapis.com \
    cloudresourcemanager.googleapis.com \
    iamcredentials.googleapis.com \
    sts.googleapis.com \
    --project=articulate-run-306102
##################################################
# Create a new workload identity pool
##################################################
gcloud iam workload-identity-pools create my-workload-pool-github \
    --location="global" \
    --description="My workload pool for github pipeline" \
    --display-name="Pool for github pipeline"
##################################################
# Create a workload identity pool provider
##################################################
gcloud iam workload-identity-pools providers create-oidc "github-provider" \
  --project="articulate-run-306102" \
  --location="global" \
  --workload-identity-pool="my-workload-pool-github" \
  --display-name="GitHub Actions Provider" \
  --issuer-uri="https://token.actions.githubusercontent.com" \
  --attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository" \
  --attribute-condition="assertion.repository == 'thuyltm/home'"
# Verify what exact condition is protecting your Workload Identity Pool
gcloud iam workload-identity-pools providers describe "github-provider" \
  --workload-identity-pool="my-workload-pool-github" \
  --location="global" \
  --format="value(attributeCondition)"

# sub: "repo:thuyltm@23312895/home@1039342672:ref:refs/heads/main"
#'projects/206108938560/locations/global/workloadIdentityPools/my-workload-pool-github/providers/github-provider'
#principal://iam.googleapis.com/projects/PROJECT_NUMBER/locations/global/workloadIdentityPools/POOL_ID/subject/SUBJECT
#principal://iam.googleapis.com/projects/206108938560/locations/global/workloadIdentityPools/my-workload-pool-github/subject/repo:thuyltm@23312895/home@1039342672:ref:refs/heads/main

