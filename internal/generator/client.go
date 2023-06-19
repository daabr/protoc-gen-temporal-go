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
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	workerpb "github.com/daabr/protoc-gen-temporal-go/proto/temporal"
)

const (
	interfaceSuffix = "TemporalClient"

	deprecationComment = "// Deprecated: Do not use."
)

func GenerateClient(g *protogen.GeneratedFile, service *protogen.Service) {
	interfaceName := service.GoName + interfaceSuffix
	exportedInterface(g, service, interfaceName)

	// Private structure.
	structName := unexport(interfaceName)
	g.P("type ", structName, " struct {")
	g.P("t ", clientPackage.Ident("Client"))
	g.P("}")
	g.P()

	// Client constructor.
	serviceComments(g, service)
	g.P("func New", interfaceName, "(c ", clientPackage.Ident("Client"), ") *", interfaceName, " {")
	g.P("return &", structName, "{c}")
	g.P("}")
	g.P()

	// Helper methods for executing workflows and activities.
	for _, method := range service.Methods {
		if isWorkflow(method) {
			startWorkflow(g, method, structName, service.GoName)
			executeWorkflow(g, method, structName, service.GoName)

			startChildWorkflow(g, method, structName, service.GoName)
			executeChildWorkflow(g, method, structName, service.GoName)
		}
		// TODO: else (activity)
	}
}

func exportedInterface(g *protogen.GeneratedFile, service *protogen.Service, interfaceName string) {
	serviceComments(g, service)

	g.Annotate(interfaceName, service.Location)
	g.P("type ", interfaceName, " interface {", service.Comments.Trailing)
	for _, m := range service.Methods {
		methodSignature(g, m, interfaceName)
	}
	g.P("}")
	g.P()
}

func serviceComments(g *protogen.GeneratedFile, service *protogen.Service) {
	// for _, comment := range service.Comments.LeadingDetached {
	// 	g.P(comment)
	// }
	g.P(strings.TrimSpace(service.Comments.Leading.String()))
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
}

func methodSignature(g *protogen.GeneratedFile, method *protogen.Method, interfaceName string) {
	methodComment(g, method, nil)

	p := contextPackage
	if isWorkflow(method) {
		p = workflowPackage
	}

	ctx := g.QualifiedGoIdent(p.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := g.QualifiedGoIdent(method.Output.GoIdent)

	g.Annotate(interfaceName+"."+method.GoName, method.Location)
	s := fmt.Sprintf("%s(ctx %s, in *%s) (*%s, error)", method.GoName, ctx, in, out)
	g.P(s, method.Comments.Trailing)
}

func methodComment(g *protogen.GeneratedFile, method *protogen.Method, suffix []string) {
	g.P(strings.TrimSpace(method.Comments.Leading.String()))

	if suffix != nil {
		g.P("//")
		for _, line := range suffix {
			g.P("// " + line)
		}
	}

	if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
}

func isWorkflow(method *protogen.Method) bool {
	w := proto.GetExtension(method.Desc.Options(), workerpb.E_Workflow).(*workerpb.Workflow)
	return w != nil
}

func unexport(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func executePrefix(g *protogen.GeneratedFile, method *protogen.Method, comment []string, structName, action, serviceName, ctx, in, out string) {
	methodComment(g, method, comment)

	s := "func (c *%s) %s%s%s(ctx %s, in *%s) %s {"
	s = fmt.Sprintf(s, structName, action, serviceName, method.GoName, ctx, in, out)
	g.P(s, method.Comments.Trailing)
}
