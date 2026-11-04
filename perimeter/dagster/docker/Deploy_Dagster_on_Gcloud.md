To deploy Dagster cheaply on Google Cloud Platform, use Dagster Open Source combined with serverless or low-cost compute components
#### Choose the Right Architecture
1. Dagster Core (OSS): Free under the Apache 2.0 license
2. Compute for Webserver & Daemon: Instead of an always-on Google Compute Engine (GCE) VM or a heavy GKE (Google Kubernetes Engine) cluster, run lightweight components on Google Cloud Run or a Free Tier small e2-micro GCE
3. Metadata & Run Storage: Run PostgreSQL inside a self-managed Free Tier Compute Engine VM
4. I/O Manager: Use Google Cloud Storage for intermediate data, assets and file
#### Cheap Deployment Pattern
1. Option A: Google Cloud Run

Package your Dagster code, webserver, and daemon into a Docker container and deploy it to Google Cloud Run

2. Option B: Single Budget Compute Engine (GCE) VM

Spin up a single small e2-micro or e2-small instance on Google Compute Engie running Ubuntu/Docker