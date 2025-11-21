load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "fortune_lib",
    srcs = ["fortune.go"],
    importpath = "home/phase1/fortune",
    visibility = ["//visibility:public"],
)

go_test(
    name = "fortune_test",
    srcs = ["fortune_test.go"],
    embed = [":fortune"],
)
