#! /bin/sh
dg plus create ci-api-token --description 'Used in home GitHub Actions' | gh secret set DAGSTER_CLOUD_API_TOKEN