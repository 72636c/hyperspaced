#!/bin/bash

set -euxo pipefail

go test -cover -v ./...

gofmt -d -s .
test -z "$(gofmt -d -s .)"

golint ./...
test -z "$(golint ./...)"

go vet ./...
go vet -vettool="$(command -v shadow)" ./...
