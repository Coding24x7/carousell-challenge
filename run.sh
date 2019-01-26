#!/bin/sh

set -x
set -e

path=${GOPATH}/src/github.com/Coding24x7/twitter-clone


cd ${path}

go build -o twitter-clone

./twitter-clone
