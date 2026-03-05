load("@crates//:defs.bzl", "all_crate_deps")
load("@rules_rust//rust:defs.bzl", "rust_binary")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "static_data_files",
    srcs = glob(["**/*"]),  # Includes all files and subdirectories
)

rust_binary(
    name = "bin",
    srcs = glob([
        "src/*.rs",
    ]),
    crate_root = "src/main.rs",
    data = [":static_data_files"],
    visibility = ["//visibility:public"],
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
)
