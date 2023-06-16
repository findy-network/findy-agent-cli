#!/bin/bash

if [[ -z "$1" ]]; then
	echo 'usage: $0 <agent-name1> <agent-name2> ...'
	exit 1
fi

play_path=${play_path:-""}

for b in "$@"; do
	a=$(basename "$b")
	agent_dir="$play_path""$a/"
	wallet_dir="$HOME/.indy_client/wallet/$a/"
	worker_dir="$HOME/.indy_client/worker/$a""_worker/"

	echo "$agent_dir" "$wallet_dir" "$worker_dir"
	echo -----------------------------------------

	rm -r "$agent_dir" "$wallet_dir" "$worker_dir"
done
