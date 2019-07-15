// Copyright 2019 The gVisor Authors.
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

package ext

import (
	"fmt"
	"os"
	"path"
	"testing"

	"gvisor.dev/gvisor/pkg/sentry/context"
	"gvisor.dev/gvisor/pkg/sentry/context/contexttest"
	"gvisor.dev/gvisor/pkg/sentry/vfs"
	"gvisor.dev/gvisor/runsc/test/testutil"
)

const (
	assetsDir = "pkg/sentry/fs/ext4/assets"
)

var (
	tinyImagePath = path.Join(assetsDir, "tiny.ext4")
)

func beginning(_ int64) int64 {
	return 0
}

func middle(i int64) int64 {
	return i / 2
}

func end(i int64) int64 {
	return i
}

// setUp creates a new MountNamespace from the given ext4 image. If err is non-nil,
// then it also returns a tearDown function which must be called at the end of the test.
func setUp(t *testing.T) (context.Context, *vfs.Filesystem, *vfs.Dentry, func(), error) {
	imagePath, err := testutil.FindFile(tinyImagePath)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("testutil.FindFile: %v", err)
	}

	f, err := os.Open(imagePath)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Mount the ext4 fs and retrieve the Inode structure for the file.
	mockCtx := contexttest.Context(t)
	fs, d, err := FilesystemType{}.NewFilesystem(mockCtx, nil, imagePath, vfs.NewFilesystemOptions{InternalData: f.Fd()})
	if err != nil {
		f.Close()
		return nil, nil, nil, nil, err
	}

	tearDown := func() {
		if err := f.Close(); err != nil {
			t.Fatalf("tearDown failed: %v", err)
		}
	}
	return mockCtx, fs, d, tearDown, nil
}

// TODO(b/134676337): Add tests.
func TestXxx(t *testing.T) {}
