### OpenID Connect provider
An OpenID Connect (OIDC) provider is an authorization server that verifies a user's identity, handles authentication, and issues signed ID tokens and access tokens to client application

[Guide How to configure GCP to trust GitHub's OIDC as a federated identity](https://docs.github.com/en/actions/how-tos/secure-your-work/security-harden-deployments/oidc-in-google-cloud-platform)

gcloud iam workload-identity-pools providers create-oidc "my-oidc-provider" \
    --location="global" \
    --workload-identity-pool="my-workload-pool" \
    --display-name="My OIDC Provider" \
    --issuer-uri="https://token.actions.githubusercontent.com" \
    --attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository"
