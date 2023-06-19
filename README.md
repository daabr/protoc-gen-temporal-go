# Go Code Generator for Temporal Workers

`protoc-gen-temporal-go` is a [protoc](https://protobuf.dev/reference/go/go-generated/)
plugin which generates Go language bindings for [Temporal](https://temporal.io/)
workers with their workflows and activities, based on service definitions in
[protocol-buffer](https://protobuf.dev/programming-guides/proto3/) files.

This methodology enables easier and safer usage of Temporal: it enforces
correctness and consistency within and across workers, intorduces best
practices seamlessly, reduces manually-written boilerplate, and improves
documentation and discoverability for developers and users.

Inspiration and background:

* [Public talk](https://www.youtube.com/watch?v=LxgkAoTSI8Q&t=680s) by [Jacob LeGrone](https://github.com/jlegrone)
  from [Datadog](https://www.datadoghq.com/) in Replay 2022
* [Another talk](https://www.youtube.com/watch?v=yeoawVIn060) by [Drew Hoskins](https://github.com/drewhoskins-stripe)
  from [Stripe](https://stripe.com/) in Replay 2022

FYI:

* This project is unrelated to <https://github.com/temporalio/api-go>, despite
  some keyword similarity in the project descriptions
* There's [another open-source implementation](https://github.com/lucasclerissepro/protoc-gen-temporal)
  for comparison, but it's incomplete
