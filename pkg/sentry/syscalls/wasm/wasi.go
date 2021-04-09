// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wasm

import (
	//	"time"
	"fmt"

	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/syserror"
	//	"gvisor.dev/gvisor/pkg/wasmvm"
	//"gvisor.dev/gvisor/pkg/log"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
	"gvisor.dev/gvisor/pkg/usermem"
)

// Full wasm syscall interface list is here: https://docs.rs/wasi/0.10.2+wasi-snapshot-preview1/wasi/wasi_snapshot_preview1/index.html
// FdWrite implements wasi_unstable/fd_write.
func FdWrite(t *kernel.Task, fd int32, src usermem.IOSequence, written int32) int32 {
	fmt.Printf("BBLU GOT WASI FDWRITE....\n")
	// call linux.writev()
	fmt.Printf("BBLU CLOCK 222:%v...\n", t.Kernel().MonotonicClock().Now())

	//	n, err := syscalls.WWritev(t, file, usermem.IOSequence{})
	//fmt.Printf("BBLU GOT %v, %v..\n", n, err)

	file := t.GetFile(fd)
	if file == nil {
		return 0
	}
	defer file.DecRef(t)

	//      src := usermem.IOSequence{}

	n, err := file.Writev(t, src)
	if err != syserror.ErrWouldBlock || file.Flags().NonBlocking {
		if n > 0 {
			// Queue notification if we wrote anything.
			file.Dirent.InotifyEvent(linux.IN_MODIFY, 0)
		}
		return 0
	}

	return int32(n)
}
