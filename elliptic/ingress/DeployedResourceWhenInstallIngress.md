```sh
helm status nginx-ingress 
```
The resources will be deployed during the nginx-ingress package installation
```sh
NAME: nginx-ingress
LAST DEPLOYED: Fri May 15 12:16:36 2026
NAMESPACE: default
STATUS: deployed
REVISION: 1
DESCRIPTION: Install complete
RESOURCES:
==> v1/ServiceAccount
NAME                          AGE
nginx-ingress-ingress-nginx   7m57s

==> v1/ClusterRole
NAME                          CREATED AT
nginx-ingress-ingress-nginx   2026-05-15T05:17:04Z

==> v1/RoleBinding
NAME                          ROLE                               AGE
nginx-ingress-ingress-nginx   Role/nginx-ingress-ingress-nginx   7m57s

==> v1/Service
NAME                                               TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
nginx-ingress-ingress-nginx-controller-admission   ClusterIP   10.103.106.175   <none>        443/TCP   7m57s
nginx-ingress-ingress-nginx-controller   LoadBalancer   10.111.107.44   10.111.107.44   80:30519/TCP,443:31222/TCP   7m57s

==> v1/Deployment
NAME                                     READY   UP-TO-DATE   AVAILABLE   AGE
nginx-ingress-ingress-nginx-controller   1/1     1            1           7m57s

==> v1/ValidatingWebhookConfiguration
NAME                                    WEBHOOKS   AGE
nginx-ingress-ingress-nginx-admission   1          7m57s

==> v1/ConfigMap
NAME                                     DATA   AGE
nginx-ingress-ingress-nginx-controller   0      7m57s

==> v1/ClusterRoleBinding
NAME                          ROLE                                      AGE
nginx-ingress-ingress-nginx   ClusterRole/nginx-ingress-ingress-nginx   7m57s

==> v1/Role
NAME                          CREATED AT
nginx-ingress-ingress-nginx   2026-05-15T05:17:04Z

==> v1/Pod(related)
NAME                                                      READY   STATUS    RESTARTS   AGE
nginx-ingress-ingress-nginx-controller-6fc4c9ff49-wfkst   1/1     Running   0          7m57s

==> v1/IngressClass
NAME    CONTROLLER             PARAMETERS   AGE
nginx   k8s.io/ingress-nginx   <none>       7m57s


TEST SUITE: None
NOTES:
The ingress-nginx controller has been installed.
It may take a few minutes for the load balancer IP to be available.
You can watch the status by running 'kubectl get service --namespace default nginx-ingress-ingress-nginx-controller --output wide --watch'

An example Ingress that makes use of the controller:
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: example
    namespace: foo
  spec:
    ingressClassName: nginx
    rules:
      - host: www.example.com
        http:
          paths:
            - pathType: Prefix
              backend:
                service:
                  name: exampleService
                  port:
                    number: 80
              path: /
    # This section is only required if TLS is to be enabled for the Ingress
    tls:
      - hosts:
        - www.example.com
        secretName: example-tls

If TLS is enabled for the Ingress, a Secret containing the certificate and key must also be provided:

  apiVersion: v1
  kind: Secret
  metadata:
    name: example-tls
    namespace: foo
  data:
    tls.crt: <base64 encoded cert>
    tls.key: <base64 encoded key>
  type: kubernetes.io/tls
```