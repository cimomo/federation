package(default_visibility = ["//visibility:public"])

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
        "//vendor/k8s.io/kubernetes/pkg/util/async:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/bandwidth:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/config:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/configz:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/dbus:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/ebtables:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/env:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/file:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/filesystem:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/flock:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/goroutinemap:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/hash:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/initsystem:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/interrupt:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/io:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/ipconfig:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/iptables:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/ipvs:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/keymutex:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/labels:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/limitwriter:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/maps:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/metrics:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/mount:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/net:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/netsh:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/node:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/nsenter:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/oom:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/parsers:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/pointer:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/procfs:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/reflector/prometheus:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/removeall:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/resourcecontainer:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/rlimit:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/selinux:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/slice:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/strings:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/sysctl:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/system:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/tail:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/taints:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/template:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/term:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/threading:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/tolerations:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/version:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/util/workqueue/prometheus:all-srcs",
    ],
    tags = ["automanaged"],
)

sh_test(
    name = "verify-util-pkg",
    size = "small",
    srcs = ["verify-util-pkg.sh"],
    data = glob(["*.go"]),
)
