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

	"google.golang.org/protobuf/compiler/protogen"
)

func startWorkflow(g *protogen.GeneratedFile, method *protogen.Method, structName, serviceName string) {
	comment := []string{
		"This method starts the workflow with pre-configured options, and returns a",
		"WorkflowRun to interact with it until completion. For more information, see",
		"https://docs.temporal.io/dev-guide/go/foundations#start-workflow-execution.",
	}
	ctx := g.QualifiedGoIdent(contextPackage.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := fmt.Sprintf("(%s, error)", g.QualifiedGoIdent(clientPackage.Ident("WorkflowRun")))

	executePrefix(g, method, comment, structName, "StartWorkflow", serviceName, ctx, in, out)

	g.P("opts := ", clientPackage.Ident("StartWorkflowOptions"), "{")
	nonDefaultStartWorkflowOptions(g, method)
	g.P("}")

	g.P("return ", "c.t.ExecuteWorkflow", "(ctx, opts, c.", method.GoName, ", in)")
	g.P("}")
	g.P()
}

func executeWorkflow(g *protogen.GeneratedFile, method *protogen.Method, structName, serviceName string) {
	comment := []string{
		"This method executes the workflow with pre-configured options, blocks until",
		"completion, and returns the output/error results. For more information, see",
		"https://docs.temporal.io/dev-guide/go/foundations#start-workflow-execution.",
	}
	ctx := g.QualifiedGoIdent(contextPackage.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := fmt.Sprintf("(*%s, error)", g.QualifiedGoIdent(method.Output.GoIdent))

	executePrefix(g, method, comment, structName, "ExecuteWorkflow", serviceName, ctx, in, out)

	g.P("opts := ", clientPackage.Ident("StartWorkflowOptions"), "{")
	nonDefaultStartWorkflowOptions(g, method)
	g.P("}")

	g.P("run, err := ", "c.t.ExecuteWorkflow", "(ctx, opts, c.", method.GoName, ", in)")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")

	g.P("var out *", g.QualifiedGoIdent(method.Output.GoIdent))
	g.P("err = run.Get(ctx, &out)")
	g.P("return out, err")
	g.P("}")
	g.P()
}

func startChildWorkflow(g *protogen.GeneratedFile, method *protogen.Method, structName, serviceName string) {
	comment := []string{
		"This method starts the workflow (as a child) with pre-configured options,",
		"and returns a Future to interact with it until completion. For more info,",
		"see https://docs.temporal.io/dev-guide/go/foundations#start-workflow-execution",
		"and https://docs.temporal.io/workflows#child-workflow.",
	}
	ctx := g.QualifiedGoIdent(workflowPackage.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := g.QualifiedGoIdent(workflowPackage.Ident("ChildWorkflowFuture"))

	executePrefix(g, method, comment, structName, "StartChildWorkflow", serviceName, ctx, in, out)

	g.P("ctx = ", workflowPackage.Ident("WithChildOptions"), "(ctx, ", workflowPackage.Ident("ChildWorkflowOptions"), "{")
	nonDefaultChildWorkflowOptions(g, method)
	g.P("})")

	g.P("return ", workflowPackage.Ident("ExecuteChildWorkflow"), "(ctx, c.", method.GoName, ", in)")
	g.P("}")
	g.P()
}

func executeChildWorkflow(g *protogen.GeneratedFile, method *protogen.Method, structName, serviceName string) {
	comment := []string{
		"This method executes the workflow (as a child) with pre-configured options,",
		"blocks until completion, and returns the output/error. For more information,",
		"see https://docs.temporal.io/dev-guide/go/foundations#start-workflow-execution",
		"and https://docs.temporal.io/workflows#child-workflow.",
	}
	ctx := g.QualifiedGoIdent(workflowPackage.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := fmt.Sprintf("(*%s, error)", g.QualifiedGoIdent(method.Output.GoIdent))

	executePrefix(g, method, comment, structName, "ExecuteChildWorkflow", serviceName, ctx, in, out)

	g.P("ctx = ", workflowPackage.Ident("WithChildOptions"), "(ctx, ", workflowPackage.Ident("ChildWorkflowOptions"), "{")
	nonDefaultChildWorkflowOptions(g, method)
	g.P("})")

	g.P("var out *", g.QualifiedGoIdent(method.Output.GoIdent))
	g.P("err := ", workflowPackage.Ident("ExecuteChildWorkflow"), "(ctx, c.", method.GoName, ", in).Get(ctx, &out)")
	g.P("return out, err")
	g.P("}")
	g.P()
}

func nonDefaultStartWorkflowOptions(g *protogen.GeneratedFile, method *protogen.Method) {
	// TODO
}

func nonDefaultChildWorkflowOptions(g *protogen.GeneratedFile, method *protogen.Method) {
	// TODO
}
