#!/bin/bash

set -e

if [[ -z "$1" ]]; then
	echo 'usage: $0 <agent-name1> <agent-name2> ...'
	exit 1
fi

for a in "$@"; do
	agent_dir="./play/$a/"

	mkdir -p "$agent_dir"
	pushd "$agent_dir" > /dev/null
	ln -s ../../run/* .
	popd > /dev/null

	"$agent_dir"register && "$agent_dir"login
done
