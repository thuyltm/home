#! /bin/bash
# 1. Login argocd first
kubectl port-forward svc/argocd-server -n argocd 8080:443
argocd login localhost:8080 \
  --username admin \
  --password argocd@123 \
  --insecure
argocd cluster add minikube --in-cluster
argocd cluster list
#SERVER                          NAME      VERSION  STATUS      MESSAGE  PROJECT
#https://kubernetes.default.svc  minikube  v1.35.1  Successful
# 2. Install
argocd app create guestbook \
  --repo https://github.com/thuyltm/home.git \
  --path spiral/argocd/guestbook \
  --dest-server https://kubernetes.default.svc \
  --dest-namespace default \
  --sync-policy automated \
  --auto-prune \
  --self-heal

# Manually sync
argocd app sync guestbook
# Watch live sync status
argocd app wait guestbook --sync