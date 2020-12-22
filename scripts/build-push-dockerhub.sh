#!/bin/bash

set -e

docker build -t fay:latest .
docker tag fay:latest starptech/fay:latest
docker push starptech/fay:latest

echo "Success"