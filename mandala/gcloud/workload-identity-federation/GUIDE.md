### Why?
Applications running outside Google Cloud can use service account keys to access Google Cloud resources. However, service account keys are powerful credentials, and can present a security risk if they are not managed correctly. 

With Workload Identity Federation, you can use Identity and Access Management (IAM) to grant IAM roles to principals that are based on federated identiites in a workload identity pool
- You can grant access to the principals on specific Google Cloud resources
- You can grant access to a service account, which can then access Google Cloud resources
A workload identity pool is an entity that lets you manage external identities

[How to use Workload Identity Federation with X.509 certificates](https://docs.cloud.google.com/iam/docs/workload-identity-federation-with-x509-certificates#console)
1. Enable the IAM, Resource Manager, Service Account Credentials and Security Token Service APIs
