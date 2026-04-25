load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "cmdv3_lib",
    srcs = ["tinybit.go"],
    importpath = "home/pacman/network/cmdv3",
    visibility = ["//visibility:private"],
    deps = [
        "//pacman/network/node",
        "@com_github_sirupsen_logrus//:logrus",
        "@com_github_spf13_cobra//:cobra",
    ],
)

go_binary(
    name = "cmdv3",
    embed = [":cmdv3_lib"],
    visibility = ["//visibility:public"],
)
