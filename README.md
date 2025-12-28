# Bazel
[Specifying external dependencies](https://github.com/bazel-contrib/rules_go/blob/master/docs/go/core/bzlmod.md#external-dependencies)<br/>


it is recommended to manage Go dependencies via go.mod. The go_deps extension parses this file directly, so external tooling such as gazelle update-repos is no longer needed.<br/>

(re-)generate BUILD files. <br/>

```sh 
bazel run //:gazelle 
```

An initial go.mod file can be created via

```sh
bazel run @rules_go//go mod init home
```

A dependency can be added via

```sh
bazel run @rules_go//go get github.com/labstack/echo/v4
```

please add these lines into MODULE.bazel [GUILINE](https://stackoverflow.com/questions/78983979/bzlmod-golang-missing-package-and-incorrect-function) in order for gazzelie automatically adds dependencies using in the BUILD file
```sh
go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")

# All *direct* Go dependencies of the module have to be listed explicitly.
use_repo(
    go_deps,
    "com_github_alexeidt_aio",
    ...
)
```

Next Run

```sh
bazel run @rules_go//go mod tidy
```

Run Target (Example)

```sh
bazel run //phase1/echo:echo
```

[Include Other External Dependencies](https://github.com/aspect-build/bazel-examples/blob/main/MODULE.bazel)
