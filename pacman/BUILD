load("@rules_go//go:def.bzl", "go_library", "go_binary")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "blockchain_lib",
    srcs = [
        "main.go"
    ],
    deps = [
        "//pacman/cli"
    ],
)

go_binary(
    name = "main",
    embed = [":blockchain_lib"],
)
