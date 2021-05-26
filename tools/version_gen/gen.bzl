_VERSION_GEN_BIN = "//tools/version_gen:bin"

def _impl(ctx):
    ctx.actions.run(
        progress_message = "Creating version file...",
        outputs = [ctx.outputs.out],
        tools = [ctx.executable._version_gen_bin],
        executable = ctx.executable._version_gen_bin,
        arguments = ["-out", ctx.outputs.out.path],
    )

    return [DefaultInfo(files = depset([ctx.outputs.out]))]

version_gen = rule(
    implementation = _impl,
    attrs = {
        "out": attr.output(mandatory = True),
        "_version_gen_bin": attr.label(
            executable = True,
            cfg = "host",
            allow_single_file = True,
            default = Label(_VERSION_GEN_BIN),
        ),
    },
)
