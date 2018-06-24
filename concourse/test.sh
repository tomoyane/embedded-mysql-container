#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

mv repository /go/src/repository

dep ensure
