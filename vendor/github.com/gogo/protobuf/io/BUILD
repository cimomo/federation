load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "full.go",
        "io.go",
        "uint32.go",
        "varint.go",
    ],
    importpath = "github.com/gogo/protobuf/io",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gogo/protobuf/proto:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["io_test.go"],
    importpath = "github.com/gogo/protobuf/io_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/gogo/protobuf/test:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
