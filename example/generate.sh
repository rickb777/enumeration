#!/bin/bash -e
cd "$(dirname $0)"/..
PATH=$PWD/bin:$PATH

function v
{
  echo "$@"
  "$@"
}

type enumeration

v rm -f example/*_enum.go
v go generate ./example
v go test ./example