package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "controller.go",
        "crudtester.go",
        "federation.go",
        "util.go",
    ],
    importpath = "k8s.io/federation/test/integration/framework",
    deps = [
        "//apis/federation/v1beta1:go_default_library",
        "//client/clientset_generated/federation_clientset:go_default_library",
        "//cmd/federation-apiserver/app:go_default_library",
        "//cmd/federation-apiserver/app/options:go_default_library",
        "//pkg/federatedtypes:go_default_library",
        "//pkg/federatedtypes/crudtester:go_default_library",
        "//pkg/federation-controller/cluster:go_default_library",
        "//pkg/federation-controller/sync:go_default_library",
        "//vendor/github.com/pborman/uuid:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/master:go_default_library",
        "//vendor/k8s.io/kubernetes/test/e2e_node/services:go_default_library",
        "//vendor/k8s.io/kubernetes/test/integration/framework:go_default_library",
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
)
