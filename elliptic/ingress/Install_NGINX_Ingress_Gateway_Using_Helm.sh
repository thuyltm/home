#! /bin/bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install my-ingress-nginx ingress-nginx/ingress-nginx \
  --namespace ingress-basic \
  --create-namespace
