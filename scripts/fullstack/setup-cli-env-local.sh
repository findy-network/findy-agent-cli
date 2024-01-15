#!/bin/bash

# echo $_ $0

if [[ $_ == $0 ]]; then
	printf 'WARNING:\tnot sourced, wont work!\n'
	printf "Usage:\t\tsource $0 [NO_TLS]\n\n"
fi

dont_use_tls=${1:-""}
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

if [[ "$dont_use_tls" == "NO_TLS" ]]; then
	echo "No TLS is used, unsetting FCLI_TLS_PATH"
	unset FCLI_TLS_PATH
else
	export FCLI_TLS_PATH="$GOPATH/src/github.com/findy-network/findy-agent/grpc/cert"
fi
