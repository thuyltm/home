load("@rules_go//go:def.bzl", "go_library", "go_binary")

go_library(
    name = "decentralizestorage_lib",
    srcs = [
        "cli.go",
        "cli_createblockchain.go",
        "cli_createwallet.go",
        "cli_getbalance.go",
        "cli_listaddress.go",
        "cli_printchain.go",
        "cli_reindexutxo.go",
        "cli_send.go",
        "cli_startnode.go",
        "main.go",
        "server.go",
    ],
    importpath = "home/pacman/decentralizestorage",
    visibility = ["//visibility:public"],
    deps = [
        "//pacman/cache",
        "//pacman/cipher",
        "//pacman/wallet",
    ],
)

go_binary(
    name = "main",
    embed = [":decentralizestorage_lib"],
)
