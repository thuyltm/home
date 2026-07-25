#! /bin/bash
# dev mode in-memory storage backend
vault server -dev -dev-root-token-id root -dev-tls
vault login
# Token (will be hidden): root
vault policy write admin admin-policy.hcl
vault token create -policy="admin"
#Key                  Value
#---                  -----
#token                hvs.CAESIAw7Hi4l5lXm8lJmm1FLwx5gaHGCk7TQTCnX_Sn399inGh4KHGh2cy4zWVhWTXhzRDFCMWFwTWhaY1ZLV0piZ1E
#token_accessor       q7H2wNJornqZSV1XilfhNQUq
#token_duration       768h
#token_renewable      true
#token_policies       ["admin" "default"]
#identity_policies    []
#policies             ["admin" "default"]
vault login
# Token (will be hidden): hvs.CAESIAw7Hi4l5lXm8lJmm1FLwx5gaHGCk7TQTCnX_Sn399inGh4KHGh2cy4zWVhWTXhzRDFCMWFwTWhaY1ZLV0piZ1E
# revoke the original root token
vault token revoke <YOUR_ROOT_TOKEN_ID>