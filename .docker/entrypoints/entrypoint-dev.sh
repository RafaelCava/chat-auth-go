#!/bin/bash

go mod tidy

go fmt ./...

mkdir -p tmp/main

swag init -g ./main/main.go -o ./main/docs

air