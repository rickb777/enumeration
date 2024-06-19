#!/bin/bash -e
cd "$(dirname $0)"

D=$(dirname $PWD)
D=$(dirname $D)

export PATH=$D/bin:$PATH
#type enumeration

function v
{
  echo "$@"
  "$@"
}

v rm -f *_enum.go simple/*_enum.go
v go generate . ./simple
