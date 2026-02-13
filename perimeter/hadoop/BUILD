load("@rules_java//java:defs.bzl", "java_binary")
load("@rules_jvm_external//:defs.bzl", "artifact")

java_binary(
    name = "WordCount",
    srcs = ["WordCount.java"],
    main_class = "WordCount",
    deps = [
        artifact("org.apache.hadoop:hadoop-common:3.4.2"),
        artifact("org.apache.hadoop:hadoop-mapreduce-client-core:3.4.2"),
        artifact("org.apache.hadoop:hadoop-hdfs-client:3.4.2"),
    ],
)


java_binary(
    name = "WordCount2",
    srcs = ["WordCount2.java"],
    main_class = "WordCount2",
    deps = [
        artifact("org.apache.hadoop:hadoop-common:3.4.2"),
        artifact("org.apache.hadoop:hadoop-mapreduce-client-core:3.4.2"),
        artifact("org.apache.hadoop:hadoop-hdfs-client:3.4.2"),
    ],
)
