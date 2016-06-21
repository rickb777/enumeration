#!/bin/bash -e
go install .

cd example
enumeration -type Base -v
enumeration -type Day -v

