#!/bin/bash -ex
cd "$(dirname "$0")"
go install tool
mage build coverage
grep -v '_enum.go' report.out | grep -v '^total:'