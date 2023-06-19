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

func startActivity(g *protogen.GeneratedFile, method *protogen.Method, structName, serviceName string) {
	comment := []string{
		"This method starts the activity with pre-configured options, and returns a",
		"Future to interact with it until completion. For more information, see",
		"https://docs.temporal.io/dev-guide/go/foundations#activity-execution.",
	}
	ctx := g.QualifiedGoIdent(workflowPackage.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := g.QualifiedGoIdent(workflowPackage.Ident("Future"))

	executePrefix(g, method, comment, structName, "StartActivity", serviceName, ctx, in, out)

	g.P("ctx = ", workflowPackage.Ident("WithActivityOptions"), "(ctx, ", workflowPackage.Ident("ActivityOptions"), "{")
	nonDefaultActivityOptions(g, method)
	g.P("})")

	g.P("return ", workflowPackage.Ident("ExecuteActivity"), "(ctx, c.", method.GoName, ", in)")
	g.P("}")
	g.P()
}

func executeActivity(g *protogen.GeneratedFile, method *protogen.Method, structName, serviceName string) {
	comment := []string{
		"This method executes the activity with pre-configured options, blocks until",
		"completion, and returns the output/error results. For more information, see",
		"https://docs.temporal.io/dev-guide/go/foundations#activity-execution.",
	}
	ctx := g.QualifiedGoIdent(workflowPackage.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := fmt.Sprintf("(*%s, error)", g.QualifiedGoIdent(method.Output.GoIdent))

	executePrefix(g, method, comment, structName, "ExecuteActivity", serviceName, ctx, in, out)

	g.P("ctx = ", workflowPackage.Ident("WithActivityOptions"), "(ctx, ", workflowPackage.Ident("ActivityOptions"), "{")
	nonDefaultActivityOptions(g, method)
	g.P("})")

	g.P("var out *", g.QualifiedGoIdent(method.Output.GoIdent))
	g.P("err := ", workflowPackage.Ident("ExecuteActivity"), "(ctx, c.", method.GoName, ", in).Get(ctx, &out)")
	g.P("return out, err")
	g.P("}")
	g.P()
}

func startLocalActivity(g *protogen.GeneratedFile, method *protogen.Method, structName, serviceName string) {
	comment := []string{
		"This method starts the activity (locally) with pre-configured options, and",
		"returns a Future to interact with it until completion. For more information,",
		"see https://docs.temporal.io/dev-guide/go/foundations#activity-execution",
		"and https://docs.temporal.io/activities#local-activity.",
	}
	ctx := g.QualifiedGoIdent(workflowPackage.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := g.QualifiedGoIdent(workflowPackage.Ident("Future"))

	executePrefix(g, method, comment, structName, "StartLocalActivity", serviceName, ctx, in, out)

	g.P("ctx = ", workflowPackage.Ident("WithLocalActivityOptions"), "(ctx, ", workflowPackage.Ident("LocalActivityOptions"), "{")
	nonDefaultLocalActivityOptions(g, method)
	g.P("})")

	g.P("return ", workflowPackage.Ident("ExecuteActivity"), "(ctx, c.", method.GoName, ", in)")
	g.P("}")
	g.P()
}

func executeLocalActivity(g *protogen.GeneratedFile, method *protogen.Method, structName, serviceName string) {
	comment := []string{
		"This method executes the activity (locally) with pre-configured options,",
		"blocks until completion, and returns the output/error. For more information,",
		"see https://docs.temporal.io/dev-guide/go/foundations#activity-execution",
		"and https://docs.temporal.io/activities#local-activity.",
	}
	ctx := g.QualifiedGoIdent(workflowPackage.Ident("Context"))
	in := g.QualifiedGoIdent(method.Input.GoIdent)
	out := fmt.Sprintf("(*%s, error)", g.QualifiedGoIdent(method.Output.GoIdent))

	executePrefix(g, method, comment, structName, "ExecuteLocalActivity", serviceName, ctx, in, out)

	g.P("ctx = ", workflowPackage.Ident("WithLocalActivityOptions"), "(ctx, ", workflowPackage.Ident("LocalActivityOptions"), "{")
	nonDefaultLocalActivityOptions(g, method)
	g.P("})")

	g.P("var out *", g.QualifiedGoIdent(method.Output.GoIdent))
	g.P("err := ", workflowPackage.Ident("ExecuteLocalActivity"), "(ctx, c.", method.GoName, ", in).Get(ctx, &out)")
	g.P("return out, err")
	g.P("}")
	g.P()
}

func nonDefaultActivityOptions(g *protogen.GeneratedFile, method *protogen.Method) {
	// TODO
}

func nonDefaultLocalActivityOptions(g *protogen.GeneratedFile, method *protogen.Method) {
	// TODO
}
