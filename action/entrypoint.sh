#!/bin/sh -l

set -e

mkdir -p /go/src/github.com/findy-network

mv findy-wrapper-go /go/src/github.com/findy-network
mv findy-agent /go/src/github.com/findy-network
mv findy-agent-cli /go/src/github.com/findy-network

cd /go/src/github.com/findy-network/findy-agent-cli

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
