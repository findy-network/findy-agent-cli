#!/bin/bash

if [[ -z "$1" ]]; then
	echo 'usage: $0 <agent-name>'
	exit 1
fi

agent_dir="./play/$1/"

rm -vr "$agent_dir"
