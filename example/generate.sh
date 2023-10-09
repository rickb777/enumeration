#!/bin/bash -e
cd "$(dirname $0)"

function v
{
  echo "$@"
  "$@"
}

type enumeration

v rm -f *_enum.go
v go generate .
