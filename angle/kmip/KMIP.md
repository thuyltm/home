### KIMP
The OASIS Key Management Interoperability Protocol (KIMP) is a widely adopted protocol for handling cryptographic workloads and secrets management for enterprise infrastructure

When an organization has services and applications that need to perform cryptographic operations, it often __delegates the key management task to an external provider via the KIMP protocol__

### Implementation
If you are looking for an open-source Key Management Interoperability Protocol (KMIP) server to run as a Docker container

[Cosmian KMS Docker](https://github.com/Cosmian/kms/pkgs/container/kms)

If you need an enterprise-grade, highly scalable open-source KMIP 2.1 server, the Cosmian KMS is an exceptional modern alternative. It provides official, highly polished images with a bundled web UI for managing cryptographic keys

```sh
docker run -p 9998:9998 --name kms ghcr.io/cosmian/kms:latest
#CONTAINER ID   IMAGE                        COMMAND                  CREATED         STATUS         PORTS                                                   NAMES
#311b3d03e5d5   ghcr.io/cosmian/kms:latest   "/nix/store/vwfzf2f5…"   9 minutes ago   Up 9 minutes   5696/tcp, 0.0.0.0:9998->9998/tcp, [::]:9998->9998/tcp   kms
```
[Guide](https://docs.cosmian.com/key_management_system/kmip_support/introduction/)

The Eviden KMS server implements both KMIP 1.x and 2.x interfaces. KMIP (Key Management Interoperability Protocol) is an OASIS standard designed to standadize communication between key management systems and encryption clients

**Connection Options**
- Binary Protocol: Available on port 5696
    - TLS secured
    - Client certificate required for authentication
- Json Protocol: Available on port 9990 via REST POST
    - Optional TLS security
    - Multiple authentication mechanisms supported
    - Endpoints:
        - /kmip: Handles KMIP 1.x and 2.x RequestMessage
        - /kmip/2_1: Sepcifially for KMIP 2.1
[Swagger UI](http://localhost:9998/swagger)

**KMIP Support**
- [Message Guide](http://localhost:9998/swagger)
The easiest way to call the KMIP API is to use the cosmian CLI or one of the Evident cloudproof libraries which provide wrapper calls in the corresponding language

Without the use of a library, the client must build the JSON TTLV messsages and issue an HTTP POST call to the /kmip/2_1 endpoint of the server

The esiest way to build JSON TTLV messages by using the KMS CLI in JSON mode to print the corresponding request and response messages

For sample JSON TTLV messages
```sh
# the CLI show the JSON TTLV requests and response
ckms --kms-print-json sym keys create --tag myKey
```
- [Authentication Guide](https://docs.cosmian.com/key_management_system/configuration/authentication/)
The KMS server supports three primary authentication methods:
1. TLS Client Certificates: Authentication based on X.509 client certificates
```sh
# the server extracts the username from the certificate's Subject Common Name (CN) field
docker run -p 9998:9998 --name kms ghcr.io/cosmian/kms-fips:latest \
    --tls-cert-file server.crt \
    --tls-key-file server.key \
    --clients-ca-cert-file client_ca.cert.pem
```
2. JWT Tokens: Authentication with OpenID-compliant JWT access tokens
```sh
# The JWT authentication provider configuration uses the format "JWT_ISSUER_URI,JWKS_URI,JWT_AUDIENCE_1,JWT_AUDIENCE_2,..."
docker run -p 9998:9998 --name kms ghcr.io/cosmian/kms:latest \
    --jwt-auth-provider="https://accounts.google.com,https://www.googleapis.com/oauth2/v3/certs,cosmian_kms"
```
3. API Tokens: Authentication using a pre-shared API token
```sh
# Generate a symmetric key and note its ID
ckms sym keys create
# Export the key in base64 format
ckms sym keys export -k <SYMMETRIC_KEY_ID> -f base64 api_token.base64
# Start the server
docker run -p 9998:9998 --name kms ghcr.io/cosmian/kms:latest \
    --api-token-id <SYMMETRIC_KEY_ID>
```
- [Authorization Guide](https://docs.cosmian.com/key_management_system/configuration/authorization/)

Every cryptographic object has an assigned owner. As an owner, a user holds the privilege to carry out all supported KMIP operations on their objects. Owners can grant access rights, allowing users to perform certain KMIP operations on an object

for example
```sh
# grant encrypt and decrypt to user "alice@example.com"
ckms access-rights grant alice@example.com encrypt decrypt
```

##### Typical workflow: per-user keys with limited permissions
Alice authenticates to the KMS (via her client certificate, JWT token, or API token) and calls the encrypt/decrypt endpoints referencing the key UID. The server verifies she holds the encrypt/decrypt permission before proceeding

```sh
# The admin creates a 256-bit AES key and tags it for easy lookup
ckms sym keys create --algorithm aes --number-of-bits 256 --tag user-alice-key
# Grant only encrypt and decrypt to alice
ckms access-rights grant alice@example.com -i [USER-ALICE-KEY] encrypt decrypt
```