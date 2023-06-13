# Go Code Generator for Temporal Workers

`protoc-gen-temporal-go` is a [protoc](https://protobuf.dev/reference/go/go-generated/)
plugin which generates Go language bindings for [Temporal](https://temporal.io/)
workers with their workflows and activities, based on service definitions in
[proto3](https://protobuf.dev/programming-guides/proto3/) protocol-buffer
files.

This methodology enables easier and safer usage of Temporal: it enforces
correctness and consistency, intorduces best practices, reduces manual
boilerplate, and improves documentation and discoverability.

Inspiration and background: a [public talk](https://www.youtube.com/watch?v=LxgkAoTSI8Q&t=633s)
by [Jacob LeGrone](https://github.com/jlegrone) from [Datadog](https://www.datadoghq.com/).

FYI:

* This project is unrelated to <https://github.com/temporalio/api-go>, despite
  some keyword similarity in the project descriptions
* There's [another open-source implementation](https://github.com/lucasclerissepro/protoc-gen-temporal)
  for comparison, but it's incomplete
