#!/bin/bash

set -e

go test -v -vet=off -race -coverprofile coverage.txt -covermode atomic ./...

echo "Success"