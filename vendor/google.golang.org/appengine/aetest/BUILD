load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "instance.go",
        "instance_vm.go",
        "user.go",
    ],
    importpath = "google.golang.org/appengine/aetest",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/appengine:go_default_library",
        "//vendor/google.golang.org/appengine/internal:go_default_library",
        "//vendor/google.golang.org/appengine/user:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["instance_test.go"],
    importpath = "google.golang.org/appengine/aetest",
    library = ":go_default_library",
    deps = [
        "//vendor/google.golang.org/appengine:go_default_library",
        "//vendor/google.golang.org/appengine/datastore:go_default_library",
        "//vendor/google.golang.org/appengine/internal:go_default_library",
        "//vendor/google.golang.org/appengine/memcache:go_default_library",
        "//vendor/google.golang.org/appengine/user:go_default_library",
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
