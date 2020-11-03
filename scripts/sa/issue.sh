#!/bin/bash

if [[ "$1" == "" && "$FCLI_SEND_CONNECTION_ID" == "" ]]; then
	echo 'Usage: $0 <conn_id or FCLI_SEND_CONNECTION_ID>'
	exit 1
fi

cmdl="go run ../../."
conn_id=${1:-$FCLI_SEND_CONNECTION_ID}
echo "connection id:"
echo "$conn_id"

# issue
ATTRS=$(eval "$cmdl" issue -i "$conn_id" --cred-def-id $CRED_ID)

echo "using these attrs:"
echo $ATTRS

# proof
eval "$cmdl" proof -i "$conn_id" --attrs "$ATTRS"

unset conn_id
unset cmdl

