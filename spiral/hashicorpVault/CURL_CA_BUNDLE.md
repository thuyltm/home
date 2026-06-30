```sh
export VAULT_CACERT='/tmp/vault-tls2295685568/vault-ca.pem'
export CURL_CA_BUNDLE=$VAULT_CACERT
```
CURL_CA_BUNDLE is an environment variable used by curl to locate the Certificate Authority (CA) bundle files required to verify SSL/TLS certificates