#!/bin/zsh

go test -run ^TestMain$ github.com/daabr/protoc-gen-temporal-go/cmd/protoc-gen-temporal-go --regenerate
