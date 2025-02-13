#!/bin/bash -e
cd "$(dirname $0)"
unset GOPATH

mkdir -p bin
export PATH=$PWD/bin:$PATH

DATE=$(date '+%F')

if [ -d .git ]; then
  VERSION=$(git describe --tags --always --dirty 2>/dev/null)
else
  VERSION=dev
fi

function v
{
  echo "$@"
  "$@"
}

v go mod download

v go build -o bin/enumeration -ldflags "-s -X main.version=$VERSION -X main.date=$DATE" .

v go clean -testcache

v rm -f internal/test/*_enum.go

v go test ./internal/parse

v ./internal/test/generate.sh

v ./example/generate.sh

v gofmt -l -w -s *.go

sleep 0.25s # wait for the files to be stable

mkdir -p temp/example # used in ./enumeration_test.go
cp example/*.go temp/example
rm -f temp/example/*_enum.go temp/example/*_test.go

v go test ./...
