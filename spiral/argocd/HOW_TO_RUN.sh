#! /bin/bash
argocd cluster add minikube --in-cluster
argocd cluster list
#SERVER                          NAME      VERSION  STATUS      MESSAGE  PROJECT
#https://kubernetes.default.svc  minikube  v1.35.1  Successful
argocd app create guestbook \
  --repo https://github.com/thuyltm/home.git \
  --path spiral/argocd \
  --dest-server https://kubernetes.default.svc \
  --dest-namespace default \
  --sync-policy automated \
  --auto-prune \
  --self-heal

# Manually sync
argocd app sync guestbook
# Watch live sync status
argocd app wait guestbook --sync