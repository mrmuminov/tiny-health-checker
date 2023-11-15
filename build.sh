#!/bin/bash
rm -rf build
env GOOS=linux GOARCH=amd64 go build -o build/linux_amd64 main.go
