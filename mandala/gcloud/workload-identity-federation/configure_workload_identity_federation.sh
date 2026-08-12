#! /bin/sh
# Run the following command to update your CLI environment
gcloud components update
# To create a new workload identity pool, execute the following command
gcloud iam workload-identity-pools create "my-workload-pool" \
  --location="global" \
  --description="My Workload Pool" \
  --display-name="Pool for external identifies"
# Prepare Trust Store
export ROOT_CERT=$(cat root.cert | sed 's/^[ ]*//g' | sed -z '$ s/\n$//' | tr '\n' $ | sed 's/\$/\\n/g')
export INTERMEDIATE_CERT=$(cat int.cert | sed 's/^[ ]*//g' | sed -z '$ s/\n$//' | tr '\n' $ | sed 's/\$/\\n/g')
# Create the Workload Identity Pool Provider
gcloud iam workload-identity-pools providers create-x509 "my-x509-provider" \
  --location="global" \
  --workload-identity-pool="my-workload-pool" \
  --display-name="My X509 Provider" \
  --trust-store-config-path="trust_store.yaml" \
  --attribute-mapping="google.subject=assertion.sub"
# Get Project Number
gcloud projects list --format="table(name, projectId, projectNumber)"

gcloud iam workload-identity-pools providers describe my-x509-provider \
    --workload-identity-pool="my-workload-pool" \
    --location="global" 
gcloud iam workload-identity-pools -h
gcloud iam workload-identity-pools providers update-x509 "my-x509-provider" \
  --location="global" \
  --workload-identity-pool="my-workload-pool" \
  --display-name="My X509 Provider" \
  --trust-store-config-path="trust_store.yaml" \
  --attribute-mapping="google.subject=assertion.subject.dn.cn"


gcloud iam workload-identity-pools create-cred-config \
  projects/206108938560/locations/global/workloadIdentityPools/my-workload-pool/providers/my-x509-provider \
    --credential-cert-path=leaf.cert \
    --credential-cert-private-key-path=leaf.key \
    --output-file=cred.json

gcloud auth login --cred-file=cred.json
#Authenticated with external account credentials for: [principal://iam.googleapis.com/projects/206108938560/locations/global/workloadIdentityPools/my-workload-pool/subject/example].
#Your current project is [articulate-run-306102]

export CLIENT_CERT_KEY=$(cat leaf.key | sed 's/^[ ]*//g' | sed -z '$ s/\n$//' | tr '\n' $ | sed 's/\$/\\n/g')
export CLIENT_CERT=$(cat leaf.cert | sed 's/^[ ]*//g' | sed -z '$ s/\n$//' | tr '\n' $ | sed 's/\$/\\n/g')

export LEAF_CERT=$(openssl x509 -in leaf.cert -out leaf.der -outform DER && cat leaf.der | openssl enc -base64 -A)
export INTERMEDIATE_CERT=$(openssl x509 -in int.cert -out int.der -outform DER && cat int.der | openssl enc -base64 -A)
export TRUST_CHAIN="[\\\"${LEAF_CERT}\\\", \\\"${INTERMEDIATE_CERT}\\\"]"


curl --key leaf.key \
--cert leaf.cert \
--request POST 'https://sts.mtls.googleapis.com/v1/token' \
--header "Content-Type: application/json" \
--data-raw '{
    "subject_token_type": "urn:ietf:params:oauth:token-type:mtls",
    "grant_type": "urn:ietf:params:oauth:grant-type:token-exchange",
    "audience": "//iam.googleapis.com/projects/206108938560/locations/global/workloadIdentityPools/my-workload-pool/providers/my-x509-provider",
    "requested_token_type": "urn:ietf:params:oauth:token-type:access_token",
    "scope": "https://www.googleapis.com/auth/cloud-platform",
    "subject_token": "'"$TRUST_CHAIN"'"
}'
#{
#  "access_token": "ya29.d.c0AZ4bNpYHFvDcF57BtvEHiok6BdsKTSFniNPrZ_kebd0-KQ-2ugkR6jWJzvBoAiPgnGsHW-evmLa-Ntvrn88kVhH1SLWsBuxfcXDxXBvHt_rld0OiH1DohGC83k_MofFRgc09KBcl5ozWZb6VWCQ2Bkp8LG2UWtkA6TjqIFkydybuN9fld1H84INri5kku0Nvtn26k6w39xA1yPyUhYoicTehSeLaRxUmOHwCRCo2CWB9b0Ah75hm2wreGyEHuq1d5IAk8tA09dKViZ9DAhwA6K5LfR951oRLbho-xjGF12JGMf2bN9DytD0XYKal5x69-4vzCwMc3CA2v_qazsKF9I5aZScM6oPu_zxVYsd1wv2Q8QRuSr-BuN8ewvy2hyZZgmTe-TD-qCiE4B6cAzSvEMf7pCEthZE8_0FGc72L3OjZ56ceWOKJKPPzeY2oEced7PEJcZ_bfEKFaS_uC9yqzMxy4o1Suhl2nDsLmGCBNH2qFsnPbpgCoSBpBYsSXVNtg39vZRYxgN6zlilwjM5hT6ykzOnouSho2kTx4SLzrEz1XcbsmbELo1Bg0zNdadaEUkx-B_U2AI2t2bxuXQJUTbkvNoqIYAqLFc-osBg7t53l3w0UTmfRha1Kbe-uCvpBfFH8G5kbkPeND6MaLjYc_6wM5yLEYRSkGVhJhYbdwPbjxxxftGwC0x2L9clB7c4esYEyBWLEHiqW8uLLtGrS2XhPF9Y1fKFEzUwGQUURRovJaytljzjSxmwFNA2-YnBijdqcUU5GrOIToC5iy_FEezu80f6X-dOl6Cm03WZfAkjlNNnCau9X5R1khHMNoJEB5nvoE70-0gtTQVp6vTpPNcE1b1-Op5i3ZgLceZz7F6I-Q9GLbbE7mAdfNhgeHVVCJqBGAvxX-WBlaS5ZgGU2MhnAZzuBBfe1UuNjxqphDdd7YmDoSQKT5_upyozaRXJavsBECvkoup0PWy9yscRqULkpacCY3z4EXDo",
#  "issued_token_type": "urn:ietf:params:oauth:token-type:access_token",
#  "token_type": "Bearer",
#  "expires_in": 3599
#}
export ACCESS_TOKEN=$(curl --key leaf.key \
--cert leaf.cert \
--request POST 'https://sts.mtls.googleapis.com/v1/token' \
--header "Content-Type: application/json" \
--data-raw '{
    "subject_token_type": "urn:ietf:params:oauth:token-type:mtls",
    "grant_type": "urn:ietf:params:oauth:grant-type:token-exchange",
    "audience": "//iam.googleapis.com/projects/206108938560/locations/global/workloadIdentityPools/my-workload-pool/providers/my-x509-provider",
    "requested_token_type": "urn:ietf:params:oauth:token-type:access_token",
    "scope": "https://www.googleapis.com/auth/cloud-platform",
    "subject_token": "'"$TRUST_CHAIN"'"
}' | jq -r '.access_token')
# Grant the Role roles/Storage Object Viewer and roles/Storage Bucket Viewer for bucket named bucket-quickstart_articulate-run-30610
curl -X GET 'https://storage.googleapis.com/storage/v1/b/bucket-quickstart_articulate-run-306102/o/example.txt?alt=media' -H "Authorization: Bearer $ACCESS_TOKEN"

gcloud storage buckets describe gs://bucket-quickstart_articulate-run-306102 --format="yam l(iamConfiguration)"

gcloud storage cp gs://bucket-quickstart_articulate-run-306102/example.txt example.txt

gcloud storage cp example.txt gs://bucket-quickstart_articulate-run-306102/



