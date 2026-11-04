The error occurs
```sh
google.auth.exceptions.MutualTLSChannelError: [Errno 2] No such file or directory: 'leaf.cert' dagster
```
because Google's authentication libraries (google-auth) are automatically attempting to establish a __Mutal TLS (mTLS)__ connection, but the required client certificate file (leaf.cert) cannot be found on your system

To explicitly tell the Google authentication helper not to look for a client certificate, set the following environment variable in the environment
```sh
GOOGLE_API_USE_CLIENT_CERTIFICATE=false
```
