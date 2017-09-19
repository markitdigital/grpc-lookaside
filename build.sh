#!/bin/bash

echo 'Cleaning'
rm -rf ./bin/*

echo 'Installing dependencies'
dep ensure

echo 'Running tests'
if ! go test -cover; then
	echo 'Unit testing failed, exiting...'
	exit 2
fi

echo 'Building for linux_amd64...'
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/linux_amd64/grpc-lookaside main.go

echo 'Building for windows_amd64...'
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/windows_amd64/grpc-lookaside.exe main.go