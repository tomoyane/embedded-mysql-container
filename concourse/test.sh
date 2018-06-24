#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

ls -la

cd repository

ls -la

dep ensure
