#!/bin/bash

go mod tidy

mkdir -p tmp/main

# CGO_ENABLED=0 swag init -g ./main/main.go -o ./main/docs

air