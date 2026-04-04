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