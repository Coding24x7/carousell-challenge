#!/bin/sh

set -x
set -e

path=${GOPATH}/src/github.com/Coding24X7/carousell-challenge


cd ${path}

go build -o carousell-challenge

./carousell-challenge
