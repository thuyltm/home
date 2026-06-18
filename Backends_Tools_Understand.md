1> Skaffold vs Helm

Skaffold and Helm are oftern used together to create a seamless development workflow.

- Helm is a package manager. It focuses on templating Kubernetes manifest, managing version
- Skaffold is a workflow automation tool. It handles the "inner loop" of development: watching for code change, rebuilding images, pushing them to a registry, and then deploying them to a cluster

2> Skaffold vs Terraform

Skaffold and Terraform are not competitors; they solve different parts of the DevOps lifecycle. 
- Terraform is used to build the "house" (infrastructure)
- Skaffold manage the building, pushing and deploying of the actual application containers.

3> ArgoCD and Jenkins

ArgoCD and Jenkins solve different parts of the software delivery lifecycle
- The primary role of Jenkins is Continuous Integration (CI). Jenkins builds the source code, runs unit tests, run vulnerability scans, and builds your Docker images
- The primary role of ArgoCD is Continuous Deployment (CD). ArgoCD detects this configuration change in Git and automatically synchronizes the update into your Kubernetes cluster

4> Terraform vs Spinnaker
- Terraform is used to provision and manage the underlying infrastructure (like servers, databases, and networks)
- Spinnaker is a CI platform that automates the deployment of your applications onto that infrastructure

5> Spinnaker vs ArgoCD
ArgoCD and Spinnaker are both powerful continuous delivery (CD) tools
- ArgoCD is a lightweight, Kubernetes-native GitOps tool. It cannot deploy to non-Kubernetes environments
- Spinnaker is an enterprise-grade, pipeline-driven platform for complex, multi-cloud deployment orchestration

6> Terraform vs Helm Charts
- Terraform is an Infrasture as Code (IaC) tool used to provision underlying cloud resources
- Helm is a Kubernetes package manager specifically built to deploy applications inside an existing cluster
"Terraform builds the house, while Helm moves in the furniture"

7> Istio vs Envoy
- Envoy is a high performance network proxy (the data plane)
- Istio is a service mesh platform that orchestrates and configures a fleet of Envoy proxies (the control plane)

8> Gateway vs Service Mesh
- API Gateway: The Front Door acts as the centralized entry point for all incoming traffic from external users, devices, or frontend clients

    Examples: AWS API Gateway, Kong, NGINX, Envoy

- Service Mesh: The Internal Highway is a dedicated infrastructure layer that manages the complex, high volume communication between different microservices running inside your environment

    Examples: Istio, Linkerd, Consul

9> Istio vs Linkerd

Istio and Linkerd are both powerful service meshes for Kubernetes
- Istio is the feature-rich, enterprise-grade powerhouse designed for highly complex, large-scale deployments

Traffic Management: Circuit breakers, fault injection, complex routing, rate limiter, multi-cluster
- Linkerd prioritizes operational simplicity, ultra-low resource usage, and ready-made ease of use

Traffic Management: Simple retries, timeouts, and traffic shifting

10> Data Plane vs Control Plane

An Istio Service Mesh abstracts network communication between microservices away from your application code into lightweight sidecar proxies (Envoy)

How it works:
- Data Plane: Consists of Envoy sidecar proxies deployed alongside every microservice. They intercept all incoming and outgoing network traffic
- Control Plane: The istiod component, which programs and manages these proxies, handling service discovery, certificate management, and configuration routing

