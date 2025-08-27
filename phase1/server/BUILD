load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "server_lib",
    srcs = ["main.go"],
    importpath = "home/phase1/server",
    visibility = ["//visibility:private"],
    deps = [
        "//phase1/server/handler",
        "//phase1/server/service",
        "@com_github_labstack_echo_v4//:go_default_library",
        "@com_github_labstack_echo_v4//middleware:go_default_library",
        "@com_github_lib_pq//:go_default_library",
    ],
)
