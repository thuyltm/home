load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "blockchain_lib",
    srcs = ["main.go"],
    importpath = "home/pacman",
    visibility = ["//visibility:private"],
    deps = ["//pacman/prototype:block_lib"],
)

go_binary(
    name = "main",
    embed = [":blockchain_lib"],
    visibility = ["//visibility:public"],
)
