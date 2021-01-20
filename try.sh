#!/bin/bash -ex
unset GOPATH
go mod download
go test -v .
go install .
gofmt -l -w -s *.go

cd example
gofmt -l -w base.go day.go month.go

enumeration -v -f -type Base -lc
enumeration -v -f -type Day
enumeration -v -f -type Month
enumeration -v -f -type Pet -unsnake -lc
enumeration -v -f -type Method -using methodStrings

cd ..
go test ./...
