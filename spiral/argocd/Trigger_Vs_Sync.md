1. Syncing

Syncing is the act of applying the configurations from your Git repo to your target cluster. It translates your desired state (YAML files) into active, running resources in Kubernetes
- Manual Sync: an administrator triggers it manually via the Argo CD UI or CLI
- Automated Sync: Handled by a Sync Policy that automatically applies changes as soon as you commit them to Git

2. Trigger
- Webhook Triggers: When you push a commit, Git sends a webhoo payload to ArgoCD, trigger the Argo CD controller to fetch the code and sync the application
- Notification Trigger: A trigger defines the conditions that will fire off an alert to Slack, Email when a sync operation changes state

### Typical Use Case: GitOps Workflow
1. Developer pushes a change to the Git repository
2. The Git provider sends a webhook to Argo CD (__the trigger__)
3. Argo CD compares the new Git state with the live cluster state
4. Argo CD updates the Kubernetes resources to match the new Git state (__the sync__)