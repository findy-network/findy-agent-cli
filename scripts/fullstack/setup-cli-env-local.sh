#!/bin/bash

GOPATH=${GOPATH:-`go env GOPATH`}

export FCLI=cli

cli=${FCLI:-findy-agent-cli}
. ../sa-compl.sh "$cli" "$cli"

if [ -f ./use-key.sh ]; then
	. ./use-key.sh
fi

if [ -z "$FCLI_KEY" ]; then
	export FCLI_KEY=`$cli new-key`
	printf "export FCLI_KEY=%s" $FCLI_KEY > use-key.sh
	echo "$FCLI_KEY" >> .key-backup
fi
export FCLI_CONFIG=./cfg.yaml
export FCLI_TLS_PATH="$GOPATH/src/github.com/findy-network/findy-agent/grpc/cert"
