#!/bin/bash -e
go install .

cd example
gofmt -l -w base.go day.go month.go

enumeration -type Base -v
enumeration -type Day -v
enumeration -type Month -v

