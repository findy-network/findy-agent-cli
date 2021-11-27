#!/bin/bash

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
export FCLI_TLS_PATH="$PWD/config/cert"

# FIDO2 server `findy-agent-auth` address
export FCLI_URL=http://localhost:8088
# Set the origin according to where our Web Wallet is hosted **important**
export FCLI_ORIGIN=http://localhost:3000

# Core agency's gRPC address
export FCLI_SERVER=localhost:50052

export FCLI=$cli
