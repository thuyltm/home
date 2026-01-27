load("@aspect_rules_py//py:defs.bzl", "py_binary", "py_image_layer", "py_library")
load("@bazel_lib//lib:transitions.bzl", "platform_transition_filegroup")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load")
load("@rules_pkg//:pkg.bzl", "pkg_tar")
load("@rules_shell//shell:sh_binary.bzl", "sh_binary")


py_library(
    name = "my_lib",
    srcs = ["my_lib.py"],
    deps = [
        "@pip//numpy",      # Dependency on numpy
        "@pip//requests",    # Dependency on requests
    ],
)

py_binary(
    name = "main",
    srcs = ["main.py", "my_lib.py", "__init__.py"],
    deps = [
        "//elliptic/locust:my_lib",
        "@pip//requests",
    ],
)

py_library(
    name = "server",
    srcs = ["server.py", "__init__.py"],
    deps = [
        "@pip//fastapi",
        "@pip//uvicorn",
    ],
)

sh_binary(
    name = "run_server",
    srcs = ["run_server.sh"],
    data = ["server.py", "__init__.py"], # Declare runtime dependencies
)

pkg_tar(
    name = "app_layer",
    srcs = ["run_server.sh"],
    mode = "0755"
)

oci_image(
    name = "image",
    # This is defined by an oci.pull() call in /MODULE.bazel
    base = "@my_fastapi_container_base",
    cmd = ["/home/phase1/fastapi/run_server.sh"], # absolute path to the script in the container
    tars = [":app_layer"],
)

# This is defined by an oci.pull() call in /MODULE.bazel

platform(
    name = "aarch64_linux",
    constraint_values = [
        "@platforms//os:linux",
        "@platforms//cpu:aarch64",
    ],
)

platform(
    name = "x86_64_linux",
    constraint_values = [
        "@platforms//os:linux",
        "@platforms//cpu:x86_64",
    ],
)

platform_transition_filegroup(
    name = "platform_image",
    srcs = [":image"],
    target_platform = select({
        "@platforms//cpu:arm64": ":aarch64_linux",
        "@platforms//cpu:x86_64": ":x86_64_linux",
    }),
)

# To build the image and load it into it into a local runtime:
# $ bazel run //phase1/fastapi:image_load
# $ docker run --rm thuyltm2201/oci_python_hello_fastapi:latest
oci_load(
    name = "image_load",
    image = ":platform_image",
    repo_tags = ["thuyltm2201/oci_python_hello_fastapi:latest"],
)

# Use this if you need a tarball:
# $ bazel build //hello_world:image_load.tar
filegroup(
    name = "image_load.tar",
    srcs = [":image_load"],
    output_group = "tarball",
)
