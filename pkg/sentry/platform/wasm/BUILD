load("//tools:defs.bzl", "go_library")

package(licenses = ["notice"])

go_library(
    name = "wasm",
    srcs = [
        "filters.go",
	"wasm.go",
        "subprocess.go",
    ],
    cgo = True,
    visibility = ["//:sandbox"],
    deps = [
 	"//pkg/sentry/wasmvm",
        "//pkg/abi/linux",
        "//pkg/context",
        "//pkg/log",
        "//pkg/procid",
        "//pkg/safecopy",
        "//pkg/seccomp",
	"//pkg/sentry/loader",
        "//pkg/sentry/arch",
        "//pkg/sentry/memmap",
        "//pkg/sentry/platform",
        "//pkg/sentry/platform/interrupt",
        "//pkg/sync",
        "//pkg/usermem",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)
