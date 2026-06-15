#! /bin/bash
kubectl apply -f secret.yaml
kubectl create secret tls caddi-tls-secret -n istio-ingress--cert=server.pem --key=server-key.pem
