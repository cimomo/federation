load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "uritemplates.go",
        "utils.go",
    ],
    importpath = "google.golang.org/api/googleapi/internal/uritemplates",
    visibility = ["//vendor/google.golang.org/api/googleapi:__subpackages__"],
)

go_test(
    name = "go_default_test",
    srcs = ["uritemplates_test.go"],
    importpath = "google.golang.org/api/googleapi/internal/uritemplates",
    library = ":go_default_library",
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
