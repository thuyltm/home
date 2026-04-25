load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "protocolv2",
    srcs = [
        "cmd.go",
        "msg.go",
        "msg_verack.go",
        "msg_version.go",
        "netaddr.go",
        "protocol.go",
    ],
    importpath = "home/pacman/network/protocolv2",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_sirupsen_logrus//:logrus",
        "//pacman/network/binary:binary"
    ],
)

go_test(
    name = "protocolv2_test",
    srcs = [
        "cmd_test.go",
        "msg_test.go",
        "msg_verack_test.go",
    ],
    embed = [":protocolv2"],
    deps = [
        "//pacman/network/binary:binary"
    ],
)
