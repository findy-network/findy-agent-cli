#!/bin/bash

if [[ $_ == $0 || $1 == "" ]]; then
	printf "Usage:\tsource %s <name>\n" "$0"

	exit 1
fi

for a in "$@"; do
	cmd="export $a=$a-`uuidgen`"
	eval "$cmd"
	printf "your variable \$%s is ready for use\n" "$a"
done
