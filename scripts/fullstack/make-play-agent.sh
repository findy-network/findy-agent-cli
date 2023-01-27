#!/bin/bash

set -e

if [[ -z "$1" ]]; then
	echo 'usage: $0 <agent-name>'
	exit 1
fi

agent_dir="./play/$1/"

mkdir -p "$agent_dir"
pushd "$agent_dir" > /dev/null
ln -s ../../run/* .
popd > /dev/null

"$agent_dir"register && "$agent_dir"login
