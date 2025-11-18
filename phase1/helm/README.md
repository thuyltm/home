# What is Helm and Why use it
Kubernetes is a powerful platform for managing containerized applications, but **deploying and maintaining these applications** can quickly become complex. That's where Helm, the open-source **package manager for Kubernetes**, comes in. 

Helm streamlines the deployment process by allowing us to define, install and manage Kubernetes applications using resuable templates known as Helm charts. A Helm chart is essentially a collection of YAML files that describe a related set of Kubernetes resources.

# Benefits of using Helm Charts
1. Helm charts allow us to **parameterize our Kubernetes manifests**. This lets us dynamically inject values--like replica counts, image tags, or resource limits--without hardcoding them in multiple places
2. Helm charts is **reusability across environments and clusters**. Whether you're spinning up your cluster in multiple environments (e.g., dev, staging, production) or deploying the same application across different clusters, Helm enables us to reuse the same chart with minimal changes.
3. Helm supports **versioned** charts

**The directory structure will be as such**
    demo-helm/
    ├── .helmignore   # Contains patterns to ignore when packaging Helm charts.
    ├── Chart.yaml    # Information about your chart
    ├── values.yaml   # The default values for your templates
    ├── charts/       # Charts that this chart depends on
    └── templates/    # The template files
       └── tests/    # The test files
       └── deployment.yaml # a basic manifest for creating a Kubernetes deployment
       └── service.yaml # a basic manifest for creating a service endpoint for your deployment
       └── _helper.tpl # a place to put template helpers that you can re-use throughout the chart

# Helm Command
```sh
helm create demo-helm
helm install NAME ./demo-helm/
helm get manifest NAME

```
When you want to test the template rendering, but not actually install anything, you can use
```sh
helm install --debug --dry-run=client test ./demo-helm/
```
But using --dry-run will make it easier to test your code, but it won't ensure that Kubernetes itself will accept the templates you generate

# Example
https://github.com/GoogleContainerTools/skaffold/tree/main/examples/helm-deployment