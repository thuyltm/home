load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "protocolv3",
    srcs = [
        "cmd.go",
        "msg.go",
        "msg_getdata.go",
        "msg_inv.go",
        "msg_pingpong.go",
        "msg_tx.go",
        "msg_verack.go",
        "msg_version.go",
        "netaddr.go",
        "protocol.go",
    ],
    importpath = "home/pacman/network/protocolv3",
    visibility = ["//visibility:public"],
    deps = [
        "//pacman/network/binary",
        "@com_github_sirupsen_logrus//:logrus",
    ],
)

go_test(
    name = "protocolv3_test",
    srcs = [
        "cmd_test.go",
        "msg_test.go",
        "msg_tx_test.go",
        "msg_verack_test.go",
    ],
    embed = [":protocolv3"],
    deps = [
        "//pacman/network/binary",
        "@com_github_google_go_cmp//cmp",
    ],
)
