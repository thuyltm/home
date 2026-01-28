package(default_visibility = ["//visibility:public"])

load("@aspect_bazel_lib//lib:transitions.bzl", "platform_transition_filegroup")
load("@rules_java//java:defs.bzl", "java_binary", "java_library", "java_test")
load("@aspect_bazel_lib//lib:tar.bzl", "tar")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load")

java_library(
    name = "java-maven-lib",
    srcs = glob(["src/main/java/com/example/myproject/*.java"]),
    deps = ["@maven//:com_google_guava_guava"],
)

java_binary(
    name = "java-maven",
    main_class = "com.example.myproject.App",
    runtime_deps = [":java-maven-lib"],
)

java_test(
    name = "tests",
    srcs = glob(["src/test/java/com/example/myproject/*.java"]),
    test_class = "com.example.myproject.TestApp",
    deps = [
        ":java-maven-lib",
        "@maven//:com_google_guava_guava",
        "@maven//:junit_junit",
    ],
)

tar(
    name = "layer",
    srcs = ["java-maven_deploy.jar"],
)

oci_image(
    name = "image",
    base = "@distroless_java",
    entrypoint = [
        "java",
        "-jar",
        "/java-maven-deploy.jar",
    ],
    tars = [":layer"],
)

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
    repo_tags = ["thuyltm2201/oci_hello_java_maven:latest"],
)
