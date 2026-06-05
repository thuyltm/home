#! /bin/bash
kubectl apply -f secret.yaml
kubectl create secret tls caddi-tls-secret -n istio-system --cert=server.pem --key=server-key.pem
