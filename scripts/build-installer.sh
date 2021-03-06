#!/bin/bash

set -e

export CGO_ENABLED=0
repro_flags="-ldflags=-buildid= -trimpath"

go fmt ./cmd/... ./internals/...
go mod tidy
go build $repro_flags -o fayinstaller cmd/installer/main.go

echo "Success"