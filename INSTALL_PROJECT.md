1. Install buildifier
```sh
sudo chmod a+x buildifier
mv buildifier /usr/local/bin
```
2. Install Bazelisk

Bazelisk is the recommended way to install Bazel on Ubuntu. It automatically downloads and installs the appropriate version of Bazel

You download Bazelisk binary on our Release page and add it to your Path manually

3. Install docker

https://docs.docker.com/engine/install/ubuntu/

Docker Hub Login Credentials store in a JSON file at ~/.docker/config.json in Linux

4. Install Bazel

https://bazel.build/install/ubuntu
```sh
sudo apt install g++ unzip zip
sudo apt-get install default-jdk
sudo apt update && sudo apt install bazel-7.7.1
cd home
bazel version
# Extracting Bazel installation...
# Starting local Bazel server and connecting to it...
# Build label: 7.7.1
# Build target: @@//src/main/java/com/google/devtools/build/lib/bazel:BazelServer
# Build time: Wed Nov 12 17:33:47 2025 (1762968827)
# Build timestamp: 1762968827
# Build timestamp as int: 1762968827
```
5. Install Go
```sh
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go[version].linux-amd64.tar.gz
```
6. Install Python

Conda lets you create virtual environment with specific Python version

https://www.anaconda.com/docs/getting-started/miniconda/install/linux-install
~~~
If you'd prefer that conda's base environment not be activated on startup, run the following command when conda is activated:

conda config --set auto_activate_base false

You can undo this later by running `conda init --reverse $SHELL`
~~~
7. Install default JDK
```sh
% sudo apt install default-jdk
% sudo update-alternatives --config java
% nano ~/.bashrc
# Scroll to the bottom and add the following lines
# export JAVA_HOME="/usr/lib/jvm/java-21-openjdk-amd64"
# export PATH="$JAVA_HOME/bin:$PATH"
% source ~/.bashrc
```
8. Install OpenSSH Server
```sh
sudo apt update
sudo apt install openssh-server
```
After installation, you must allow ssh traffic through port 22
```sh
sudo systemctl status ssh
sudo systemctl enable ssh
sudo ufw allow ssh
```
9. Install Node

https://nodejs.org/en/download
10. Install Scala

https://www.scala-lang.org/download/
~~~
Checking if ~/.local/share/coursier/bin is in PATH
  Should we add ~/.local/share/coursier/bin to your PATH via ~/.profile? [Y/n] y

Checking if the standard Scala applications are installed
  Installed ammonite
  Installed cs
  Installed coursier
  Installed scala
  Installed scalac
  Installed scala-cli
  Installed sbt
  Installed sbtn
  Installed scalafmt
~~~
11. Install Rust

https://rust-lang.org/tools/install/
~~~
Rustup metadata and toolchains will be installed into the Rustup
home directory, located at:

  /home/thuy/.rustup

This can be modified with the RUSTUP_HOME environment variable.

The Cargo home directory is located at:

  /home/thuy/.cargo

This can be modified with the CARGO_HOME environment variable.

The cargo, rustc, rustup and other commands will be added to
Cargo's bin directory, located at:

  /home/thuy/.cargo/bin

This path will then be added to your PATH environment variable by
modifying the profile files located at:

/home/thuy/.profile
/home/thuy/.bashrc

You can uninstall at any time with rustup self uninstall and
these changes will be reverted.
~~~
12. Instal pkg-config

pkg-config is a build development tool in Ubuntu used to retrieve information about installed libraries, such as compiler flags (--cflags) and linker flags (--libs)
```sh
sudo apt-get update
sudo apt-get install pkg-config build-essential
```
13. Install Kubernete

https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/

14. Install Minikube

https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download

15. Create a Personal Access Token (Classic) in your GitHub account

setup a PAT (Personal Access Token) for GitHub container registry (GHCR) to authenticate Docker operations (login, push and pull) in place of your password

Personal access tokens (classic) function like ordinary OAuth access token. They can be used instead of a password for Git over HTTPS, or can be used to authenticate to the API over Basic AUthentication

Follow the sequences of steps
```md
1. Navigate directly to [your account's GitHub Developer Settings](https://github.com/settings/apps)
2. In the left sidebar, click Personal access tokens and select Tokens (classic)
3. Clock the Generate new token dropdown and choose Generate new token (classic)
4. Check the boxes for the scopes or permissions you want to grant this token (e.g., select repo for basic Git command-line access)
5. Click Generate
6. echo $(cat PAT_for_GitHub_Account.md) | docker login ghcr.io -u YOUR_GITHUB_USERNAME --password-stdin
```

16. Install Helm
```sh
tar -zxvf helm-v4.2.0-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/helm
```

17. Install skaffold
```sh
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
sudo install skaffold /usr/local/bin/
```

18. Install Kompose
```sh
curl -L https://github.com/kubernetes/kompose/releases/download/v1.34.0/kompose-linux-amd64 -o kompose
chmod +x kompose
sudo mv ./kompose /usr/local/bin/kompose
```

19. Generate self-signed certificates

[Guide](https://docs.opensearch.org/latest/security/configuration/generate-certificates/)

20. Install Terraform
```sh
wget -O - https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(grep -oP '(?<=UBUNTU_CODENAME=).*' /etc/os-release || lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform
```

21. Install Local Kind Cluster
```sh
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.32.0/kind-linux-amd64
```

22. Install GitHub CLI

gh is GitHub on the command line. It brings pull requests, issues, and other GitHub concepts to the terminal where you are already working with git and your code
```sh
# https://cli.github.com/
out=$(mktemp) && wget -nv -O$out https://cli.github.com/packages/githubcli-archive-keyring.gpg 
cat $out | sudo tee /etc/apt/keyrings/githubcli-archive-keyring.gpg > /dev/null
sudo chmod go+r /etc/apt/keyrings/githubcli-archive-keyring.gpg
sudo mkdir -p -m 755 /etc/apt/sources.list.d
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
sudo apt update
sudo apt install gh -y
```

23. Install Google Cloud cli

[Guide](https://docs.cloud.google.com/sdk/docs/install-sdk)
```sh
curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-linux-x86_64.tar.gz
tar -xf google-cloud-cli-linux-x86_64.tar.gz
./google-cloud-sdk/install.sh
gcloud init
gcloud auth 
gcloud config list
#For example, to view the help for gcloud compute instances create:
gcloud help compute instances create
```
Create a Service Account
```md
1. Go to the Google Cloud Console Service Account page
2. Select your Google Cloud Project
3. Click Create Service Account at the top of the screen
4. Assign Roles: Select the specific roles your code needs (e.g., Storage Object Viewer, BigQuery Admin)
5. Generate Key
6. Nagivate to the Keys tab, click Add Key
```