load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["scanner.go"],
    importpath = "github.com/hashicorp/hcl/hcl/scanner",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/hashicorp/hcl/hcl/token:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["scanner_test.go"],
    importpath = "github.com/hashicorp/hcl/hcl/scanner",
    library = ":go_default_library",
    deps = ["//vendor/github.com/hashicorp/hcl/hcl/token:go_default_library"],
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
