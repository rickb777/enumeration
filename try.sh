#!/bin/bash -ex
unset GOPATH
go mod download
go test .
go install .
gofmt -l -w -s *.go

cd example
gofmt -l -w base.go day.go month.go

enumeration -type Base -v -f -lc
enumeration -type Day -v -f
enumeration -type Month -v -f
enumeration -type Pet -v -f -unsnake -lc

go test .
