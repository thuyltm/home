load("@rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "database_lib",
    srcs = [
        "block.go",
        "blockchain.go",
        "cli.go",
        "main.go",
        "proofofwork.go",
    ],
    importpath = "home/pacman/decentralizestorage",
    deps = ["@com_github_boltdb_bolt//:bolt"],
)

go_binary(
    name = "main",
    embed = [":database_lib"],
)
