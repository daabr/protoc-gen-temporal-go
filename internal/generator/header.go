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
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	Executable = "protoc-gen-temporal-go"
	Version    = "0.0.0"
)

const (
	// Package field number in "FileDescriptorProto" in
	// https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto.
	fileDescriptorProtoPackageFieldNumber = 2

	// Syntax field number in "FileDescriptorProto" in
	// https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto.
	fileDescriptorProtoSyntaxFieldNumber = 12
)

func GenerateHeader(g *protogen.GeneratedFile, f *protogen.File, ver string) {
	// Attach all comments associated with the syntax field.
	path := protoreflect.SourcePath{fileDescriptorProtoSyntaxFieldNumber}
	leadingComments(g, f.Desc.SourceLocations().ByPath(path))

	g.P(fmt.Sprintf("// Code generated by %s. DO NOT EDIT.", Executable))
	g.P("// versions:")
	g.P("// - ", Executable, " v", Version)
	alignment := len(Executable) - len("protoc") + 1
	g.P("// - protoc", strings.Repeat(" ", alignment), ver)
	if f.Proto.GetOptions().GetDeprecated() {
		g.P("// ", f.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", f.Desc.Path())
	}
	g.P()

	// Attach all comments associated with the package field.
	path = protoreflect.SourcePath{fileDescriptorProtoPackageFieldNumber}
	leadingComments(g, f.Desc.SourceLocations().ByPath(path))

	g.P("package ", f.GoPackageName)
	g.P()
}

func leadingComments(g *protogen.GeneratedFile, loc protoreflect.SourceLocation) {
	for _, s := range loc.LeadingDetachedComments {
		g.P(protogen.Comments(s))
		g.P()
	}
	if s := loc.LeadingComments; s != "" {
		g.P(protogen.Comments(s))
		g.P()
	}
}