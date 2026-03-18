load("@rules_go//go:def.bzl", "go_library", "go_binary")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "block_lib",
    srcs = [
        "block.go",
        "blockchain.go",
        "main.go"
    ],
    importpath = "home/pacman/prototype",
)

go_binary(
    name = "main",
    embed = [":block_lib"],
)

