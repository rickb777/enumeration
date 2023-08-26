#!/bin/bash -e
cd "$(dirname $0)"
unset GOPATH

mkdir -p bin
export PATH=$PWD/bin:$PATH

function v
{
  echo "$@"
  "$@"
}

v go mod download
#v go test .
v go build -o bin/enumeration .
#type enumeration

v go clean -testcache

v rm -f internal/test/*_enum.go

v go test ./internal/parse

v ./internal/test/generate.sh

v ./example/generate.sh

v gofmt -l -w -s *.go

sleep 0.25s # wait for the files to be stable

v go test ./...
