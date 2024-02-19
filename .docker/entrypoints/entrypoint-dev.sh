#!/bin/bash

go mod tidy

go fmt ./...

mkdir -p tmp/main

# CGO_ENABLED=0 swag init -g ./main/main.go -o ./main/docs

air