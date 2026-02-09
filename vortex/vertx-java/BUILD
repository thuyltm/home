load("@contrib_rules_jvm//java:defs.bzl", "java_test_suite")

#load("@contrib_rules_jvm//java:defs.bzl", "java_test_suite", "junit5_deps")
load("@rules_java//java:defs.bzl", "java_binary", "java_library", "java_test")
load("@rules_jvm_external//:defs.bzl", "artifact")

package(default_visibility = ["//visibility:public"])

java_library(
    name = "java-maven-lib",
    srcs = glob(["src/main/java/com/example/starter/*.java"]),
    deps = [
        "@maven//:com_google_guava_guava",
        "@maven//:io_vertx_vertx_core",
        "@maven//:io_vertx_vertx_launcher_application",
    ],
)

java_binary(
    name = "java-maven",
    main_class = "com.example.starter.MainVerticle",
    runtime_deps = [":java-maven-lib"],
)

java_test_suite(
    name = "othertests",
    size = "small",
    srcs = glob(["src/test/java/com/example/starter/OtherTests.java"]),
    runner = "junit5",
    test_suffixes = ["Tests.java"],
    runtime_deps = [
        artifact("org.junit.jupiter:junit-jupiter-engine"),
        artifact("org.junit.platform:junit-platform-launcher"),
        artifact("org.junit.platform:junit-platform-reporting"),
    ],
    deps = [
        artifact("org.junit.jupiter:junit-jupiter-api"),
        artifact("com.google.guava:guava"),
    ],
)

java_test_suite(
    name = "sampleVerticleTests",
    size = "small",
    srcs = glob(["src/test/java/com/example/starter/SampleVerticleTests.java"]),
    runner = "junit5",
    test_suffixes = ["Tests.java"],
    runtime_deps = [
        artifact("org.junit.jupiter:junit-jupiter-engine"),
        artifact("org.junit.platform:junit-platform-launcher"),
        artifact("org.junit.platform:junit-platform-reporting"),
    ],
    deps = [
        artifact("org.junit.jupiter:junit-jupiter-api"),
        artifact("com.google.guava:guava"),
        ":java-maven-lib",
        "@maven//:io_vertx_vertx_core",
        "@maven//:io_vertx_vertx_junit5",
        "@maven//:io_vertx_vertx_launcher_application",
        "@maven//:io_vertx_vertx_web_client",
        "@maven//:io_vertx_vertx_web_common",
    ],
)

java_test(
    name = "sampleVerticleTestsCheck",
    srcs = glob(["src/test/java/com/example/starter/SampleVerticleTests.java"]),
    # Pass arguments to the ConsoleLauncher to discover and run tests
    args = [
        "execute",
        "--select-package",
        "com.example.starter",  # Replace with your actual test package
        "--details",
        "tree",
    ],
    # Use the ConsoleLauncher as the entry point for JUnit 5
    main_class = "org.junit.platform.console.ConsoleLauncher",
    test_class = "com.example.starter.SampleVerticleTests",
    runtime_deps = [
        "@maven//:org_junit_platform_junit_platform_console",
    ],
    deps = [
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_junit_jupiter_junit_jupiter_engine",
        # Add other project dependencies here
        artifact("com.google.guava:guava"),
        ":java-maven-lib",
        "@maven//:io_vertx_vertx_core",
        "@maven//:io_vertx_vertx_junit5",
        "@maven//:io_vertx_vertx_launcher_application",
        "@maven//:io_vertx_vertx_web_client",
        "@maven//:io_vertx_vertx_web_common",
    ],
)
