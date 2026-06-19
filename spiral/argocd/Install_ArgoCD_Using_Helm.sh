#! /bin/bash

# 1. Add the official Argo Helm repository
helm repo add argo https://argoproj.github.io/argo-helm

# 2. Update your local Helm chart cache
helm repo update

# 3. Install Argo CD into a dedicated 'argocd' namespace
helm install argocd argo/argo-cd --namespace argocd --create-namespace

# 4. Update a single value quickly
helm upgrade argocd argo/argo-cd --namespace argocd --reuse-values --set server.ingress.enabled=true

# 5. Export the configuration into a local files
helm get values argocd -n argocd -o yaml > old-values.yaml
helm upgrade argocd argo/argo-cd -f old-values.yaml

