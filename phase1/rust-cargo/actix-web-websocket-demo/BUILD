package(default_visibility = ["//visibility:public"])

load("@crates//:defs.bzl", "all_crate_deps")
load("@rules_rust//rust:defs.bzl", "rust_binary")

filegroup(
    name = "static_data_files",
    srcs = glob(["**/*"]), # Includes all files and subdirectories
)

rust_binary(
    name = "bin",
    crate_root = "src/main.rs",
    srcs = glob([
        "src/*.rs",
    ]),
    data = [":static_data_files"],
    deps = [
        # External crates
        "@crates//:actix-files",
        "@crates//:actix-web",
        "@crates//:actix-ws",
        "@crates//:env_logger",
        "@crates//:futures-util",
        "@crates//:log",
        "@crates//:tokio",
    ],
    visibility = ["//visibility:public"],
)
