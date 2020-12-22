#!/bin/bash

set -e

go tool cover -html=coverage.out

echo "Success"