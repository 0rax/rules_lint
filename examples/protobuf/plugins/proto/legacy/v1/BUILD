load("@rules_proto//proto:defs.bzl", "proto_library")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "legacy_proto",
    srcs = ["legacy.proto"],
    strip_import_prefix = "/plugins/proto",
)
