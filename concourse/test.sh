#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

ls -la

mv -t repository /go/src/repository

cd /go/src/repository

ls -la

dep ensure
