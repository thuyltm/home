1. Pre-download maven packages before compiling
```shell
bazel fetch @maven//:all
```
2. To cache Maven downloads in Bazel and generate a lock file, configure the lock file __maven_install.json__ in MODULE.bazel.
```shell
bazel run @maven//:pin
```
Then generate or update the __maven_install.json__ file, run:
```shell
REPIN=1 bazel run @maven//:pin
# OR
REPIN=1 bazel sync --only=maven
```

