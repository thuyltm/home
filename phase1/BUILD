load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "hello",
    embed = [":phase1_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "phase1_lib",
    srcs = ["hello.go"],
    importpath = "home/phase1",
    visibility = ["//visibility:private"],
    deps = ["//phase1/fortune"],
)
