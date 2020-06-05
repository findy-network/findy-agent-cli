#!/bin/sh -l

set -e

mkdir -p /go/src/github.com/optechlab

mv findy-go /go/src/github.com/optechlab
mv findy-agent /go/src/github.com/optechlab
mv findy-cli /go/src/github.com/optechlab

cd /go/src/github.com/optechlab/findy-cli

echo "Install deps"
go get -t ./...

echo "Check formatting"
GOFILES=$(find . -name '*.go')
gofmt -l $GOFILES

echo "Run vet"
go vet ./...
go vet -vettool=$GOPATH/bin/shadow ./...

echo "Run lint"
golint -set_exit_status ./...

echo "::set-output name=time::$time"
