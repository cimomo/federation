package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "federation-apiserver",
    importpath = "k8s.io/federation/cmd/federation-apiserver",
    library = ":go_default_library",
)

go_library(
    name = "go_default_library",
    srcs = ["apiserver.go"],
    importpath = "k8s.io/federation/cmd/federation-apiserver",
    deps = [
        "//cmd/federation-apiserver/app:go_default_library",
        "//cmd/federation-apiserver/app/options:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/flag:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/logs:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/version/verflag:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//cmd/federation-apiserver/app:all-srcs",
    ],
    tags = ["automanaged"],
)
