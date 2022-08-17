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

v ./internal/test/generate.sh

v ./example/generate.sh

v go clean -testcache

sleep 0.25s # wait for the files to be stable

v go test ./...
