go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")

# All *direct* Go dependencies of the module have to be listed explicitly.
use_repo(
    go_deps
)