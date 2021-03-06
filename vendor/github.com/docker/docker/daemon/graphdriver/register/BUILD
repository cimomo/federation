load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "register_vfs.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "register_aufs.go",
            "register_btrfs.go",
            "register_devicemapper.go",
            "register_overlay.go",
            "register_zfs.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "register_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/daemon/graphdriver/register",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/daemon/graphdriver/vfs:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/docker/daemon/graphdriver/aufs:go_default_library",
            "//vendor/github.com/docker/docker/daemon/graphdriver/btrfs:go_default_library",
            "//vendor/github.com/docker/docker/daemon/graphdriver/devmapper:go_default_library",
            "//vendor/github.com/docker/docker/daemon/graphdriver/overlay:go_default_library",
            "//vendor/github.com/docker/docker/daemon/graphdriver/overlay2:go_default_library",
            "//vendor/github.com/docker/docker/daemon/graphdriver/zfs:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/docker/docker/daemon/graphdriver/lcow:go_default_library",
            "//vendor/github.com/docker/docker/daemon/graphdriver/windows:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
