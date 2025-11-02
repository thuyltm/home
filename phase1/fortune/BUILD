load("@rules_go//go:def.bzl", "go_library", "go_test")
load("@rules_oci//oci:defs.bzl", "oci_image")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

go_library(
    name = "fortune_lib",
    srcs = ["fortune.go"],
    importpath = "home/phase1/fortune",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "my_app",
    srcs = ["main.go"],
    deps = [":my_library"],
)

pkg_tar(
    name = "fortune_layer",
    srcs = [":fortune_bin"],
    mode = "0755",  # Executable permissions
    package_dir = "/home/thuy/Documents/Learn/home/bazel-bin/phase1/fortune",  # Or any desired path within the image
)

oci_image(
    name = "fortune_image",
    base = "@golang_base_linux_amd64",  # Or another base image, defined via oci.pull in MODULE.bazel
    entrypoint = ["/home/thuy/Documents/Learn/home/bazel-bin/phase1/fortune/fortune_bin"],
    tars = [":fortune_layer"],
)

go_test(
    name = "fortune_test",
    srcs = ["fortune_test.go"],
    embed = [":fortune"],
)
