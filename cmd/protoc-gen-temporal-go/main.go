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

// Protoc-gen-temporal-go is a [protoc] plugin which generates Go language
// bindings for [Temporal] workers, comprising their workflows and activities,
// based on [proto3] service definitions. This enforces correctness and
// consistency while reducing boilerplate.
//
// Background - a talk by [Jacob LeGrone] from [Datadog]:
// <https://www.youtube.com/watch?v=LxgkAoTSI8Q&t=633s>.
//
// [protoc]: https://protobuf.dev/reference/go/go-generated/
// [Temporal]: https://temporal.io/
// [proto3]: https://protobuf.dev/programming-guides/proto3/
// [Jacob LeGrone]: https://github.com/jlegrone
// [Datadog]: https://www.datadoghq.com/
package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/daabr/protoc-gen-temporal-go/internal/generator"
)

const filenameSuffix = "_temporal.pb.go"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("%s %s\n", generator.Executable, generator.Version)
		return
	}

	protogen.Options{}.Run(func(p *protogen.Plugin) error {
		v := protocVersion(p)
		for _, f := range p.Files {
			if !f.Generate {
				continue
			}
			generateFile(p, f, v)
		}
		return nil
	})
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	s := fmt.Sprintf("v%d.%d.%d", v.GetMajor(), v.GetMinor(), v.GetPatch())
	if v.GetSuffix() != "" {
		s += "-" + v.GetSuffix()
	}
	return s
}

func generateFile(p *protogen.Plugin, f *protogen.File, ver string) *protogen.GeneratedFile {
	if len(f.Services) == 0 {
		return nil
	}
	filename := f.GeneratedFilenamePrefix + filenameSuffix
	g := p.NewGeneratedFile(filename, f.GoImportPath)
	generator.GenerateHeader(g, f, ver)
	for _, service := range f.Services {
		generator.GenerateWorker(g, service)
	}
	return g
}
