#!/bin/bash

set -euxo pipefail

golint ./...
go test -cover -v ./...
go vet -all -shadow ./...
