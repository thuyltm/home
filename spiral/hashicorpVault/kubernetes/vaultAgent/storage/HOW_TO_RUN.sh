#! /bin/sh
#kubectl create configmap db-init-script --from-file=init.sql
#skaffold delete
#kubectl patch pvc postgres-storage-postgres-0 -p '{"metadata":{"finalizers":null}}'
skaffold deploy
kubectl exec -it postgres-0 -- psql -U myuser -d mydb
# Run a docker image as a temporary pod to test the connection to the database
kubectl run pg-test-client --image=postgres:latest --restart=Never --rm -i --tty -- psql postgresql://myuser:MyHome123@postgres-service.default.svc.cluster.local:5432/mydb
psql -U myuser -d mydb -f init-db/init.sql
psql -U myuser -d mydb -c "\dt vault_kv_store"
psql -U myuser -d mydb -c "\dt"
