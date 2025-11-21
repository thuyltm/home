[Gateway API](https://gateway-api.sigs.k8s.io/implementations/)

Implementaions generally fall into two categories, which are call profiles:
- **Gateway** controllers reconcile the Gateway resource and are intended to handle north-south traffic, mainly concerned with comming from outside the cluster to inside
- **Mesh** controllers reconcile Service resouces with HTTPRoutes attached and are intended to handle east-west traffic, within the same cluster or set of clusters
