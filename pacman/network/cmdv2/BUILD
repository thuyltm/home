load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "cmdv2_lib",
    srcs = ["tinybit.go"],
    importpath = "home/pacman/network/cmdv2",
    visibility = ["//visibility:private"],
    deps = [
        "//pacman/network/node",
        "@com_github_sirupsen_logrus//:logrus",
        "@com_github_spf13_cobra//:cobra",
    ],
)

go_binary(
    name = "cmdv2",
    embed = [":cmdv2_lib"],
    visibility = ["//visibility:public"],
)
