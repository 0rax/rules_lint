# buildifier: disable=bzl-visibility
load("@rules_lint//lint/private:machine_report_testing.bzl", "report_results_count_test", "report_test")
load("//test:machine_output.bzl", "machine_buf_modules_report", "machine_buf_plugins_report", "machine_buf_report")
load("//tools/lint:linters.bzl", "buf_modules_test", "buf_plugins_test", "buf_test")

package(default_visibility = ["//visibility:public"])

buf_test(
    name = "buf",
    srcs = ["//src:foo_proto"],
    expected_exit_code = 1,
)

machine_buf_report(
    name = "machine_buf_report",
    src = "//src:foo_proto",
)

report_test(
    name = "buf_machine_output_test",
    expected_tool = "Buf",
    expected_uri = "src/file.proto",
    report = "machine_buf_report",
)

buf_modules_test(
    name = "buf_modules",
    srcs = [
        "//modules/fixtures/widget/v1:widget_proto",
        "//modules/prefixed:gadget_proto",
        "//modules/proto/acme/v1:acme_proto",
    ],
    expected_exit_code = 1,
)

machine_buf_modules_report(
    name = "machine_buf_modules_report",
    src = "//modules/proto/acme/v1:acme_proto",
)

machine_buf_modules_report(
    name = "machine_buf_modules_fixtures_report",
    src = "//modules/fixtures/widget/v1:widget_proto",
)

machine_buf_modules_report(
    name = "machine_buf_modules_prefixed_report",
    src = "//modules/prefixed:gadget_proto",
)

report_test(
    name = "buf_modules_machine_output_test",
    expected_tool = "Buf",
    expected_uri = "modules/proto/acme/v1/greeting.proto",
    report = "machine_buf_modules_report",
)

report_results_count_test(
    name = "buf_modules_machine_output_count_test",
    expected_count = 2,
    report = "machine_buf_modules_report",
)

report_test(
    name = "buf_modules_fixtures_machine_output_test",
    expected_tool = "Buf",
    expected_uri = "modules/fixtures/widget/v1/widget.proto",
    report = "machine_buf_modules_fixtures_report",
)

report_test(
    name = "buf_modules_prefixed_machine_output_test",
    expected_tool = "Buf",
    expected_uri = "modules/prefixed/gadget.proto",
    report = "machine_buf_modules_prefixed_report",
)

buf_plugins_test(
    name = "buf_plugins",
    srcs = ["//plugins/proto/legacy/v1:legacy_proto"],
    expected_exit_code = 1,
)

machine_buf_plugins_report(
    name = "machine_buf_plugins_report",
    src = "//plugins/proto/legacy/v1:legacy_proto",
)

report_test(
    name = "buf_plugins_machine_output_test",
    expected_tool = "Buf",
    expected_uri = "plugins/proto/legacy/v1/legacy.proto",
    report = "machine_buf_plugins_report",
)

# buf.yaml enables only the plugin's rule, so the finding comes from the plugin
report_results_count_test(
    name = "buf_plugins_machine_output_count_test",
    expected_count = 1,
    report = "machine_buf_plugins_report",
)
