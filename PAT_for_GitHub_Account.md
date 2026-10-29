You cannot use your standard GitHub password to log into the container registry. You must use a Personal Access Token
1. Go to your GitHub account settings Settings -> Developer Settings -> Personal Access Tokens -> Tokens (classic).
2. Generate a new token with the read:packages scope (and write:packages if you intend to push images).

```sh
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/ 
```
This error __ERROR: failed to build: failed to solve: failed to fetch oauth token: denied: denied__ occurs when the Docker build engine (BuildKit) tries to pull the uv image from the GitHub Container Registery (ghcr.io), but its authentication handshake is reject

__ghcr.io/astral-sh/uv:latest__ is a public image, you shouldn't need credentials to access it. However, if your environment is already logged into a GitHub account, you many need to explicitly authenticate using a PAT for handling anonymously handshake

To create a Personal Access Token (PAT) on GitHub, you navigate to Developer Settings and create Fine-grained tokens

Log into the registry via your terminal
```sh
export YOUR_GITHUB_PAT=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
export YOUR_GITHUB_USERNAME=thuyltm
echo $YOUR_GITHUB_PAT | docker login ghcr.io -u thuyltm --password-stdin
```