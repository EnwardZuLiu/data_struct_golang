#!/bin/sh

go test

# test list package
cd list
go test

cd ../sort
go test