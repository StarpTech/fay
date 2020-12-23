#!/bin/bash

set -e

./scripts/lint.sh

go test -v -vet=off -race -coverprofile coverage.out -covermode atomic ./...

echo "Success"