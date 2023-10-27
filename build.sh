#!/bin/sh

GOOS=linux go build -ldflags="-w -s" -o desafio3 cmd/main.go cmd/wire.go