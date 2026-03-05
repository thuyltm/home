load("@rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "proofofwork_lib",
    srcs = [
        "block.go",
        "blockchain.go",
        "main.go",
        "proofofwork.go",
    ],
    importpath = "home/pacman/proofofwork",
)

go_binary(
    name = "main",
    embed = [":proofofwork_lib"],
)
