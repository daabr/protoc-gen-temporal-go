/*
MIT License

Copyright (c) 2023 Daniel Abraham

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"flag"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// When this environment variable is set, we skip running tests and instead
// run main(). This allows the test binary to run "protoc" as a subprocess
// and pass itself to protoc as a plugin.
const runtimeMode = "RUN_MAIN_INSTEAD_OF_TESTS"

// Use --regenerate to regenerate the golden .pb.go files.
var regenerate = flag.Bool("regenerate", false, "regenerate golden files")

func init() {
	if _, ok := os.LookupEnv(runtimeMode); ok {
		main()
		os.Exit(0)
	}
}

func TestMain(t *testing.T) {
	// Initialize temporary directory for compiled proto output files.
	workDir, err := os.MkdirTemp("", "protoc-gen-temporal-go-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(workDir)

	// Initialize list of test cases (proto input files to be compiled).
	tests := []string{}
	err = filepath.WalkDir("../../testdata", func(path string, d fs.DirEntry, err error) error {
		if !strings.HasSuffix(path, ".proto") {
			return nil
		}
		tests = append(tests, path)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// Run all test cases (compile all input proto files, compare output files
	// to pre-compiled golden pb.go files).
	for _, proto := range tests {
		name := strings.TrimSuffix(filepath.Base(proto), ".proto")
		t.Run(name, func(t *testing.T) {
			runProtoc(t, proto, workDir)
			got := readOutputFile(t, proto, workDir)
			want := readGoldenFile(t, proto)
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("content mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func runProtoc(t *testing.T, inputProtoFile, workDir string) {
	ex, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	args := []string{
		"--plugin=protoc-gen-temporal-go=" + ex,
		"--temporal-go_out=" + workDir,
		"--temporal-go_opt=paths=source_relative",
		"--fatal_warnings",
		"--proto_path=../../proto",                     // .../temporal/worker.proto
		"--proto_path=../../submodules/temporalio/api", // .../temporal/...
		"--proto_path=" + filepath.Dir(inputProtoFile),
		inputProtoFile,
	}
	cmd := exec.Command("protoc", args...)
	cmd.Env = append(os.Environ(), runtimeMode+"=1")
	t.Log("running:\n", strings.Join(cmd.Args, "\n"))
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		t.Logf("protoc output:\n%s", out)
	}
	if err != nil {
		t.Fatal("protoc error:\n", err)
	}
}

func readGoldenFile(t *testing.T, inputProtoFile string) string {
	name := strings.TrimSuffix(inputProtoFile, ".proto") + filenameSuffix
	b, err := os.ReadFile(name)
	if err != nil {
		t.Log("golden file not found: ", name)
		return ""
	}
	s := string(b)
	s = s[strings.Index(s, "package"):]
	t.Log("want:\n", s)
	return s
}

func readOutputFile(t *testing.T, inputProtoFile, workDir string) string {
	goldenName := strings.TrimSuffix(inputProtoFile, ".proto") + filenameSuffix
	name := strings.TrimSuffix(filepath.Base(inputProtoFile), ".proto") + filenameSuffix
	name = filepath.Join(workDir, name)
	b, err := os.ReadFile(name)
	if err != nil {
		t.Log("output file not found: ", name)
		if *regenerate {
			os.Remove(goldenName)
		}
		return ""
	}
	s := string(b)
	s = s[strings.Index(s, "package"):]
	t.Log("got:\n", s)
	if *regenerate {
		if err := os.WriteFile(goldenName, b, 0o644); err != nil {
			t.Error(err)
		}
	}
	return s
}
