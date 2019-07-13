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

// Binary proctor is a utility that facilitates language testing for Go112.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

var (
	list      = flag.Bool("list", false, "list all available tests")
	test      = flag.String("test", "", "run a single test from the list of available tests")
	version   = flag.Bool("v", false, "print out the version of node that is installed")
	dir       = "/usr/local/go/test"
	testRegEx = regexp.MustCompile(`^.+\.go$`)
	exclDirs  = regexp.MustCompile(`^.+\/(bench|stress|safe)\/.+$|^.+\.dir.+$`)
)

func main() {
	flag.Parse()

	if *list && *test != "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *list {
		listTests()
		return
	}
	if *version {
		fmt.Println("Go version: 1.12.5 is installed.")
		return
	}
	runTest(*test)
}

func listTests() {
	// Go tool dist test tests.
	args := []string{"tool", "dist", "test", "-list"}
	cmd := exec.Command("go", args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to list: %v", err)
	}

	// Go tests on disk.
	var tests []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		name := filepath.Base(path)

		if info.IsDir() {
			return nil
		}

		if !testRegEx.MatchString(name) {
			return nil
		}

		if exclDirs.MatchString(path) {
			return nil
		}

		tests = append(tests, path)
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to walk %q: %v", dir, err)
	}

	for _, file := range tests {
		fmt.Println(file)
	}
}

func runTest(test string) {
	toolArgs := []string{
		"tool",
		"dist",
		"test",
	}
	diskArgs := []string{
		"run",
		"run.go",
		"-v",
	}
	if test != "" {
		if _, err := os.Stat(test); !os.IsNotExist(err) {
			relPath, err := filepath.Rel(dir, test)
			if err != nil {
				log.Fatalf("Failed to get rel path: %v", err)
			}
			diskArgs = append(diskArgs, "--", relPath)
			cmd := exec.Command("go", diskArgs...)
			cmd.Dir = dir
			cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("Failed to run: %v", err)
			}
			return
		}
		toolArgs = append(toolArgs, "-run", test)
		cmd := exec.Command("go", toolArgs...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to run: %v", err)
		}
		return
	}
	runAllTool := exec.Command("go", toolArgs...)
	runAllTool.Stdout, runAllTool.Stderr = os.Stdout, os.Stderr
	if err := runAllTool.Run(); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
	runAllDisk := exec.Command("go", diskArgs...)
	runAllDisk.Dir = dir
	runAllDisk.Stdout, runAllDisk.Stderr = os.Stdout, os.Stderr
	if err := runAllDisk.Run(); err != nil {
		log.Fatalf("Failed to run disk tests: %v", err)
	}
}
