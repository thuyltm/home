#! /bin/sh
# Get Project Number
gcloud projects list --format="table(name, projectId, projectNumber)"
#NAME              PROJECT_ID             PROJECT_NUMBER
#My Project 76456  articulate-run-306102  206108938560
# identify the cloud billing account linked to a project
gcloud beta billing projects describe articulate-run-306102