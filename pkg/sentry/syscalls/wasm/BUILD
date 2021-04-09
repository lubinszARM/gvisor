load("//tools:defs.bzl", "go_library")

package(licenses = ["notice"])

go_library(
    name = "wasm",
    srcs = [
        "wasi.go",
    ],
    marshal = True,
    visibility = ["//:sandbox"],
    deps = [
        "//pkg/abi/linux",
        "//pkg/syserror",
        "//pkg/log",
	"//pkg/usermem",
	"//pkg/sentry/kernel",
	"//pkg/sentry/syscalls/linux",
    ],
)
