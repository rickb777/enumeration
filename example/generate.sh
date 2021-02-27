#!/bin/bash -e
cd "$(dirname $0)"

function v
{
  echo "$@"
  "$@"
}

v gofmt -l -w $(ls *.go | fgrep -v _)

v enumeration -v -f -type Base -lc
v enumeration -v -f -type Day
v enumeration -v -f -type Month -ic
v enumeration -v -f -type Pet -unsnake -lc -using petStrings
v enumeration -v -f -type Method -ic -using methodStrings
v enumeration -v -f -type GreekAlphabet -using greekStrings
