#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

ls -la
pwd
# mv repository /go/src/repository

cd repository

ls -la

dep ensure
