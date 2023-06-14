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

package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	workerpb "github.com/daabr/protoc-gen-temporal-go/proto/temporal"
)

func GenerateWorker(g *protogen.GeneratedFile, service *protogen.Service) {
	worker := proto.GetExtension(service.Desc.Options(), workerpb.E_Worker).(*workerpb.Worker)

	g.P("func StartWorker", service.GoName, "(c ", clientPackage.Ident("Client"), ") {")
	g.P(`taskQueue := "`, worker.TaskQueue, `"`)
	g.P("workerOptions := ", workerPackage.Ident("Options"), "{")
	nonDefaultWorkerOptions(g, worker.Options)
	g.P("}")

	g.P("w := ", workerPackage.Ident("New"), "(c, taskQueue, workerOptions)")
	g.P()

	registerWorkerMethods(g, service.Methods)
	g.P()

	g.P("if err := w.Run(", workerPackage.Ident("InterruptCh"), "()); err != nil {")
	g.P(logPackage.Ident("Fatalln"), `("Failed to start Temporal worker:", err)`)
	g.P("}")

	g.P("}")
	g.P()
}

func nonDefaultWorkerOptions(g *protogen.GeneratedFile, o *workerpb.WorkerOptions) {
	options := []struct {
		goName       string
		value        interface{}
		defaultValue interface{}
	}{
		{
			"MaxConcurrentActivityExecutionSize",
			o.MaxConcurrentActivityExecutionSize,
			0,
		},
	}
	for _, option := range options {
		if option.value != option.defaultValue {
			g.P(option.goName, ": ", option.value, ",")
		}
	}
}

func registerWorkerMethods(g *protogen.GeneratedFile, methods []*protogen.Method) {
	for _, m := range methods {
		w := proto.GetExtension(m.Desc.Options(), workerpb.E_Workflow).(*workerpb.Workflow)
		if w != nil {
			g.P("w.RegisterWorkflow(", m.GoName, ")")
		} else {
			g.P("w.RegisterActivity(", m.GoName, ")")
		}
	}
}
