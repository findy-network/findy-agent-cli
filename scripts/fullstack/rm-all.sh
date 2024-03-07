#!/bin/bash

# NOTE: we must not use this. there's wallets and worker-wallets that not exist.
# set -e

location=$(dirname "$BASH_SOURCE")
[[ $location = . ]] && location="./play" || path="."

if [[ -z "$1" ]]; then
	dirs="$location/*/"
else
	dirs="$@"
fi

for d in $dirs; do
	./rm-play-agent.sh `basename $d`
done

