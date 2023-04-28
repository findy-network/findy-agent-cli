#!/bin/bash

calc_name() {
	local debug=${1:-""}

	local location=$(dirname "$BASH_SOURCE")
	local name=$(basename "$location")
	[[ "$debug" != "" ]] && echo "name from location: $name" >&2
	[[ "$name" == "." ]] && name=$(basename "${PWD}")
	[[ "$debug" != "" ]] && echo "name from PWD: $name" >&2

	if [[ "$name" == ".." ]]; then
		pushd .. > /dev/null
		name=$(basename "${PWD}")
		[[ "$debug" != "" ]] && echo "name: $name" >&2
		popd > /dev/null
	fi
	echo -n "$name"
}
