#!/bin/bash

set -e

if [[ -z $3 ]]; then
	echo "usage: $0 <agent-dir> <from> <to> [wait_sec]"
	exit 1
fi

wait_sec=${4:-"0.1"}

for ((i = $2; i < $3; i++ ))
do
	echo "Starting to background: ""$i"
	./make-agent.sh "$1"-"$i" &
	echo "wait: ""$wait_sec"
	sleep "$wait_sec"
done

