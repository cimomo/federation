load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "k8s.io/gengo/examples/defaulter-gen",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/k8s.io/gengo/args:go_default_library",
        "//vendor/k8s.io/gengo/examples/defaulter-gen/generators:go_default_library",
    ],
)

go_binary(
    name = "defaulter-gen",
    importpath = "k8s.io/gengo/examples/defaulter-gen",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/k8s.io/gengo/examples/defaulter-gen/generators:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
