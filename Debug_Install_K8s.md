1. Debug
```sh
kubectl api-resources --verbs=list --namespaced -o name | xargs -n 1 kubectl get --show-kind --ignore-not-found -n caddi
```
2. Check your User ID
```sh
kubectl exec <pod-name> -- id
```