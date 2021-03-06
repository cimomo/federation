load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "bolt_unix.go",
        "boltsync_unix.go",
        "bucket.go",
        "cursor.go",
        "db.go",
        "doc.go",
        "errors.go",
        "freelist.go",
        "node.go",
        "page.go",
        "tx.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "bolt_amd64.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "bolt_amd64.go",
            "bolt_linux.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "bolt_amd64.go",
            "bolt_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/boltdb/bolt",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "freelist_test.go",
        "node_test.go",
        "page_test.go",
    ],
    importpath = "github.com/boltdb/bolt",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "bucket_test.go",
        "cursor_test.go",
        "db_test.go",
        "quick_test.go",
        "simulation_test.go",
        "tx_test.go",
    ],
    importpath = "github.com/boltdb/bolt_test",
    deps = [":go_default_library"],
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
        "//vendor/github.com/boltdb/bolt/cmd/bolt:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
