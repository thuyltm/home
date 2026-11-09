
### Always Free Tier (Non-Expiring)
[Link](https://cloud.google.com/free?hl=en)

Even after your trial ends or without starting a paid plan, Google Cloud provides non-expiring monthly limits across 24+ core products. Key free limits include:
- Compute Engine: 1 non-preemptible e2-micro VM instance per month in select US regions. If you run exactly one e2-micro instance continuously for the entire month, it will be 100% free because its run time will exactly match the total hours of the month.
- Cloud Storage: 5 GB of standard storage (US regions) per month. Google Cloud Storage provides __free tier allowances for Standard storage only__, while Nearline storage has no free tier
- Cloud Run: 2 million requests per month.

• --min-instances 0: Crucial. Tells Cloud Run to scale down to zero instances when no one is accessing your site. If set to 1 or more, it runs continuously and will quickly exhaust your free allocation.

• --max-instances 1: Caps the maximum container scaling. This protects you from unexpected billing spikes if your application experiences a sudden flood of automated bot traffic or a DDoS attack.
- BigQuery: 1 TB of querying and 10 GB of storage per month.
- Google Kubernetes Engine (GKE): 

__Clean Up After Practice__ The absolute best way to maintain a $0 balance is to script your environment (e.g., using a quick setup framework or Terraform). Spin up the cluster when you start studying, and delete the cluster completely when you are done for the day.

__Avoid Load Balancers__: Creating a Kubernetes Service with type: LoadBalancer automatically provisions a Google Cloud Load Balancer, which charges a fixed hourly rate regardless of traffic

- Cloud Build: 120 build-minutes per day.
- Cloud Run Functions: 2 million calls per month

### Key Cost-Saving Tips on GCP
1. Avoid Cloud SQL if possible for small loads
2. Leverage GCS Lifecycle Rules: Automatically delete or transition old temporary data, intermediate Parquet/CSV files in your GCS buckets to lower storage tiers (Nearline/Coldline)
3. Right-size compute for heavy tasks: Don't run heavy data transformations directly on the Dagster webserver VM. Offload heavy data processing to serverless jobs or trigger BigQuery SQL queries
4. Static IP vs Ephemeral IP: Leave the VM's external IP as Ephemeral
5. Network Egress: Google Cloud gibes you 1 GB of free internet data egress per month from the US region to non-China/Austrilia destinations. Keep your application hosted in the same region if possible to keep internal traffic free