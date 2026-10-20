#! /bin/sh
# Deploy an existin project to Dagster Cloud using GitHub Actions
# From your Dagster project directory with your virtual environment active
uv add dagster-cloud
dg plus login
# Comple the login process in your browser
dg plus deploy configure
# Follow the prompts to configure your deployment and
# CREATE THE GITHUB ACTION TOKEN ACCESS
#dg plus create ci-api-token --description 'Used in home GitHub Actions' | gh secret set DAGSTER_CLOUD_API_TOKEN
# Lastly, commit and push the changes to your repository to trigger the GitHub Action workflow
# get DAGSTER_CLOUD_API_TOKEN
dg plus config view --show-token
#Config file: /home/thuy/.config/dg.toml
#default_deployment: prod
#organization: dagster-dbt-home
#user_token: user:b0063a89da0b48e4a5b0b15fbb8da827