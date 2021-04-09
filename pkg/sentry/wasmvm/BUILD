load("//tools:defs.bzl", "go_library", "go_test")

package(licenses = ["notice"])

go_library(
    name = "wasmvm",
    srcs = [
        "wasmthread.go",
        "wasmproc.go",
        "wasmtimevm.go",
	"wasmvm.go",
    ],
    cgo = True,
    visibility = ["//visibility:public"],
    deps = [
	"//pkg/abi/linux",
	"//pkg/log",
	"//pkg/context",
	"//pkg/syserror",
	"//pkg/usermem",
	"//pkg/sentry/kernel",
	"//pkg/sentry/syscalls/wasm",
	"@com_github_bytecodealliance_wasmtime_go//:go_default_library",
	],
)
