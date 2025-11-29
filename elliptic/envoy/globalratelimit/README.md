1. Install Redis because of Envoy Ratelimit requires a Redis instance as its caching layer
2. Enable Global Rate liit in Envoy Gateway
```sh
helm upgrade eg oci://docker.io/envoyproxy/gateway-helm \
  --set config.envoyGateway.rateLimit.backend.type=Redis \
  --set config.envoyGateway.rateLimit.backend.redis.url="redis.redis-system.svc.cluster.local:6379" \
  --reuse-values \
  -n envoy-gateway-system
```
3. Rate Limit Distinct Users except Admin
4. Testing
```sh
for i in {1..4}; do curl -I --resolve "www.example.com:80:$( echo $GATEWAY_HOST )" --header "x-user-id: one" www.example.com/serviceb; sleep 1; done
# Output
HTTP/1.1 200 OK
date: Sat, 29 Nov 2025 04:16:16 GMT
content-length: 30
content-type: text/plain; charset=utf-8
x-ratelimit-limit: 3, 3;w=3600
x-ratelimit-remaining: 2
x-ratelimit-reset: 2624

HTTP/1.1 200 OK
date: Sat, 29 Nov 2025 04:16:17 GMT
content-length: 30
content-type: text/plain; charset=utf-8
x-ratelimit-limit: 3, 3;w=3600
x-ratelimit-remaining: 1
x-ratelimit-reset: 2623

HTTP/1.1 200 OK
date: Sat, 29 Nov 2025 04:16:18 GMT
content-length: 30
content-type: text/plain; charset=utf-8
x-ratelimit-limit: 3, 3;w=3600
x-ratelimit-remaining: 0
x-ratelimit-reset: 2622

HTTP/1.1 429 Too Many Requests
x-ratelimit-limit: 3, 3;w=3600
x-ratelimit-remaining: 0
x-ratelimit-reset: 2621
date: Sat, 29 Nov 2025 04:16:18 GMT
transfer-encoding: chunked
```

```sh
for i in {1..4}; do curl -I --resolve "www.example.com:80:$( echo $GATEWAY_HOST )" --header "x-user-id: admin" www.example.com/serviceb; sleep 1; done
# Output
HTTP/1.1 200 OK
date: Sat, 29 Nov 2025 04:17:39 GMT
content-length: 30
content-type: text/plain; charset=utf-8

HTTP/1.1 200 OK
date: Sat, 29 Nov 2025 04:17:40 GMT
content-length: 30
content-type: text/plain; charset=utf-8

HTTP/1.1 200 OK
date: Sat, 29 Nov 2025 04:17:41 GMT
content-length: 30
content-type: text/plain; charset=utf-8

HTTP/1.1 200 OK
date: Sat, 29 Nov 2025 04:17:42 GMT
content-length: 30
content-type: text/plain; charset=utf-8
```