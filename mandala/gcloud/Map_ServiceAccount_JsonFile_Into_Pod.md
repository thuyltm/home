### To map a Google Cloud service account JSON key file to a K8s pod

1. store the JSON file as a Kubernetes Secret
```sh
kubectl k create secret generic gcp-service-account-secret \
     --from-file=key.json=./perimeter/dagster/dagster-postgresql-quickstart/articulate-run-306102-910c63dbe0d6.json
```
2. Mount the secret as a volume
3. Set the GOOGLE_APPLICAITON_CREDENTIALS environment variable to point to the mounted file
```sh
spec:
  containers:
  - name: app-container
    image:
    env:
      - name: GOOGLE_APPLICATION_CREDENTIALS
        value: /var/secrets/google/key.json
    volumeMounts:
      - name: google-cloud-key
        mountPath: /var/secrets/google
        readOnly: true
  volumes:
  - name: google-cloud-key
    secret:
      secretName: gcp-service-account-secret
```