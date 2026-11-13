load("@rules_go//go:def.bzl", "go_binary")

# buf runs a plugin by its name in buf.yaml: the binary is named after it.
go_binary(
    name = "buf-plugin-message-no-legacy-prefix",
    srcs = ["main.go"],
    visibility = ["//visibility:public"],
    deps = [
        "@build_buf_go_bufplugin//check",
        "@build_buf_go_bufplugin//check/checkutil",
        "@org_golang_google_protobuf//reflect/protoreflect",
    ],
)
