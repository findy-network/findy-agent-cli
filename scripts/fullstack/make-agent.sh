#!/bin/bash

if [[ -z "$1" ]]; then
	echo 'usage: $0 <agent-name>'
	exit 1
fi

mkdir -p "$1"
pushd "$1" > /dev/null
ln -s ../run/* .
popd > /dev/null

"$1"/register && "$1"/login
