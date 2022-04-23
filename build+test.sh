#!/bin/bash -e
cd "$(dirname $0)"
unset GOPATH

function v
{
  echo "$@"
  "$@"
}

v go mod download
#v go test .
v go install .
v gofmt -l -w -s *.go

v ./example/generate.sh

v go clean -testcache

v go test ./...
