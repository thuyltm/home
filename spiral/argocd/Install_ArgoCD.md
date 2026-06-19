[Guide](https://dev.to/durrello/setting-up-argocd-on-minikube-for-a-local-dev-environment-5637)

In order to access the server UI you have the following options:

1. kubectl port-forward service/argocd-server -n argocd 8080:443

    and then open the browser on http://localhost:8080 and accept the certificate

2. 
    
    2.1. enable ingress in the values file `server.ingress.enabled` and either
      - Add the annotation for ssl passthrough: https://argo-cd.readthedocs.io/en/stable/operator-manual/ingress/#option-1-ssl-passthrough

      Refer to SSL_PASSTHROUGH_GATEWAY.yaml, 

      - Set the `configs.params."server.insecure"` in the values file and terminate SSL at your ingress: https://argo-cd.readthedocs.io/en/stable/operator-manual/ingress/#option-2-multiple-ingress-objects-and-hosts

      Consult SSL_TERMINATE_AT_GATEWAY_GRPC.yaml, SSL_TERMINATE_AT_GATEWAY_HTTP.yaml

    2.2. you must then configure the API server should be run with TLS disabled. Edit the _argocd-server_ deployment to add the _--insecure_ flag to the argocd-server command, or simple set __server.insecure: "true"__ in the __argocd-cmd-params-cm__ ConfigMap

### Update the random password generated during the installation
```sh
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
argocd login localhost:8080 \
  --username admin \
  --password <YOUR_PASSWORD_HERE> \
  --insecure
argocd account update-password \
  --current-password <YOUR_PASSWORD_HERE> \
  --new-password <NEW_STRONG_PASSWORD>
```

#### Install ArgoCD CLI
```sh
curl -sSL -o argocd-linux-amd64 https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64
sudo install -m 555 argocd-linux-amd64 /usr/local/bin/argocd
rm argocd-linux-amd64
```
