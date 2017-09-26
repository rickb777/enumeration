#!/bin/bash -ex
go get github.com/onsi/gomega
go test .
go install .
gofmt -l -w -s *.go

cd example
gofmt -l -w base.go day.go month.go

enumeration -type Base -v -f
enumeration -type Day -v -f
enumeration -type Month -v -f

go test .
