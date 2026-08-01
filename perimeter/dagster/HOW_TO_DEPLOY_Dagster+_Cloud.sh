#! /bin/sh
# Deploy an existin project to Dagster Cloud using GitHub Actions
# From your Dagster project directory with your virtual environment active
uv add dagster-cloud
dg plus login
# Comple the login process in your browser
dg plus deploy configure
# Follow the prompts to configure your deployment and
# create the GitHub Action Token Access
#dg plus create ci-api-token --description 'Used in home GitHub Actions' | gh secret set DAGSTER_CLOUD_API_TOKEN
# Lastly, commit and push the changes to your repository to trigger the GitHub Action workflow