1> Skaffold vs Helm

Skaffold and Helm are oftern used together to create a seamless development workflow.

- Helm is a package manager. It focuses on templating Kubernetes manifest, managing version
- Skaffold is a workflow automation tool. It handles the "inner loop" of development: watching for code change, rebuilding images, pushing them to a registry, and then deploying them to a cluster

2> Skaffold vs Terraform

Skaffold and Terraform are not competitors; they solve different parts of the DevOps lifecycle. Terraform is used to build the "house" (infrastructure), while Skaffold manage the building, pushing and deploying of the actual application containers.