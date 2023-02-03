#!/bin/bash

if [[ -z "$1" ]]; then
	echo 'usage: $0 <agent-name>'
	exit 1
fi

agent_dir="./play/$1/"
wallet_dir="$HOME/.indy_client/wallet/$1/"
worker_dir="$HOME/.indy_client/worker/$1_worker/"

echo "$agent_dir" "$wallet_dir" "$worker_dir"
echo -----------------------------------------

rm -r "$agent_dir" "$wallet_dir" "$worker_dir"

