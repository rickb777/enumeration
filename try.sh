#!/bin/bash -e
go install .

cd example
enumeration -type Base -i sample1.go -o out1.go

cd ..
mkdir -p example/foo
enumeration -type Base -i example/sample1.go -o example/foo/out2.go -package example

diff example/out1.go example/foo/out2.go

