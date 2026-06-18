istioctl install
istioctl install -y -f export_via_grpc.yaml
istioctl install -y -f export_via_http.yaml
istioctl install -y -f skywalking_tracing.yaml
# Enable Data Plane
kubectl label namespace book istio-injection=enabled
# If use gateway service istio-ingress/istio-ingress
kubectl get services -n istio-ingress
#NAME            TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                                      AGE
#istio-ingress   LoadBalancer   10.100.30.107   10.100.30.107   15021:30918/TCP,80:31685/TCP,443:31190/TCP   4d
curl -v  -H "Host: book.bookinfo.com" http://10.100.30.107:80/productpage | grep -o "<title>.*</title>"
# If use gateway service istio-system/istio-ingressgateway
kubectl get services -n istio-system
#NAME                          TYPE           CLUSTER-IP       EXTERNAL-IP      PORT(S)                                          AGE
#istio-ingressgateway          LoadBalancer   10.109.253.242   10.109.253.242   15021:32000/TCP,80:31305/TCP,443:31726/TCP       19h
curl -v  -H "Host: book.bookinfo.com" http://10.109.253.242:80/productpage | grep -o "<title>.*</title>"

#################################################TESTING#####################################################################
for i in $(seq 1 100); do curl -s -o /dev/null -H "Host: book.bookinfo.com" "http://$GATEWAY_URL/productpage"; done
#############################################################################################################################
##########################################VISUALIZE GRAPH####################################################################
istioctl dashboard kailik
# execute a program periodically
watch -n 1 curl -s -o /dev/null -H "Host: book.bookinfo.com" "http://$GATEWAY_URL/productpage"
##########################################SKYWALKING TRACING#################################################################
istioctl dashboard skywalking