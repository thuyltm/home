load("@rules_cc//cc:cc_binary.bzl", "cc_binary")
load("@rules_cc//cc:cc_library.bzl", "cc_library")


cc_binary(
    name = "cube",
    srcs = ["cube.c"],
    deps = [
        ":opengl_library",
    ],
)

# Use system pkg-config to find dependencies
cc_library(
    name = "opengl_library",
    copts = ["-g", "-Wall", "`pkg-config --cflags glu glut gl`"],
    linkopts = ["-lGLU -lglut -lGL"]#pkg-config --libs glu glut gl,
    # Or, preferably, use proper rules_foreign_cc to manage this
)
