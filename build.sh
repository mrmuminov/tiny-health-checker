#!/bin/bash
rm -rf build

# linux amd64
env GOOS=linux GOARCH=amd64 go build -o build/linux_amd64 main.go

# windows amd64
env GOOS=windows GOARCH=amd64 go build -o build/windows_amd64.exe main.go
