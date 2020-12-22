#!/bin/bash

set -e

go test -v -vet=off -race -coverprofile coverage.out -covermode atomic ./...

echo "Success"