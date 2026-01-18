####  Create a Self-Signed Certificate
```sh
openssl req -x509 -newkey rsa:2048 -nodes \
  -subj "/CN=localhost" \
  -keyout cert/key.pem -out cert/cert.pem -days 365
```

#### Error
failed to sufficiently increase receive buffer size (was: 208 kiB, wanted: 7168 kiB, got: 416 kiB). See https://github.com/quic-go/quic-go/wiki/UDP-Buffer-Sizes for details.

Fix: You may increase the maximum send buffer size:
```sh
sudo sysctl -w net.core.wmem_max=7500000
```

#### Testing
```sh
% curl --cacert cert/cert.pem --http3-only https://localhost:4433/
Hello from HTTP/3! You requested / via HTTP/3.0
```