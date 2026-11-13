"Define linter aspects"

load("@rules_lint//lint:buf.bzl", "lint_buf_aspect")
load("@rules_lint//lint:lint_test.bzl", "lint_test")

buf = lint_buf_aspect(
    config = Label("@//:buf.yaml"),
)

buf_test = lint_test(aspect = buf)

# an example of setting up a buf aspect for a v2 buf.yaml listing `modules`
buf_modules = lint_buf_aspect(
    config = Label("@//modules:buf.yaml"),
    infer_module = True,
)

buf_modules_test = lint_test(aspect = buf_modules)

# an example of setting up a buf aspect for a buf.yaml listing local check `plugins`
buf_plugins = lint_buf_aspect(
    config = Label("@//plugins:buf.yaml"),
    infer_module = True,
    plugins = [Label("@//plugins/buf-plugin-message-no-legacy-prefix")],
)

buf_plugins_test = lint_test(aspect = buf_plugins)
