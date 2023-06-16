#!/bin/bash

if [[ -z "$1" ]]; then
	echo 'usage: $0 <agent-name1> <agent-name2> ...'
	exit 1
fi

location=$(dirname "$BASH_SOURCE")
[[ $location = . ]] && path="play/" || path="./"
play_path="$path" "$location"/rm-agent.sh $@

