load("@rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

go_binary(
    name = "main",
    embed = [":transaction_lib"],
)

go_library(
    name = "transaction_lib",
    srcs = [
        "block.go",
        "blockchain.go",
        "cli.go",
        "main.go",
        "proofofwork.go",
        "transaction.go",
    ],
    importpath = "home/pacman/transaction",
    deps = ["@com_github_boltdb_bolt//:bolt"],
)
