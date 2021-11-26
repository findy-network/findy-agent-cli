#!/bin/bash

cli=${FCLI:-findy-agent-cli}
. ../sa-compl.sh "$cli" "$cli"

if [ -z "$FCLI_KEY" ]; then
	export FCLI_KEY=`$cli new-key`
	printf "export FCLI_KEY=%s" $FCLI_KEY > use-key.sh
	echo "$FCLI_KEY" >> .key-backup
fi
export FCLI_CONFIG=./cfg.yaml
export FCLI_TLS_PATH="$PWD/config/cert"
export FCLI_URL=http://localhost:8088
export FCLI_SERVER=localhost:50052
export FCLI_ORIGIN=http://localhost:3000
export FCLI=$cli
