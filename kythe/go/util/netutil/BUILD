load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "netutil",
    srcs = ["netutil.go"],
)

go_test(
    name = "netutil_test",
    size = "small",
    srcs = ["netutil_test.go"],
    library = "netutil",
    visibility = ["//visibility:private"],
    deps = ["//kythe/go/test/testutil"],
)
