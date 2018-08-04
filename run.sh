#!/bin/sh

set -x
set -e

path=${GOPATH}/src/github.com/Coding24x7/carousell-challenge


cd ${path}

go build -o carousell-challenge

./carousell-challenge
