To bridge your local host to kind's internal network and expose LoadBalancer services directly, use cloud-provider-kind
```sh
go install sigs.k8s.io/cloud-provider-kind@latest
```
Run the provider in a separate terminal. Keep this process running
```sh
cloud-provider-kind
```