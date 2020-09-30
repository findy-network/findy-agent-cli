#!/bin/sh -l

set -e

mkdir -p /go/src/github.com/findy-network

mv /findy-wrapper-go /go/src/github.com/findy-network
mv /findy-agent /go/src/github.com/findy-network
mv /findy-agent-cli /go/src/github.com/findy-network

cd /go/src/github.com/findy-network/findy-agent-cli

echo "Install deps"
go get -t ./...

echo "Run e2e tests"
make e2e_ci
