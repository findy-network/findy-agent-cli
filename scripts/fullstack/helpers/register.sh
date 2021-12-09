#!/bin/bash

if [[ -z $3 ]]; then
	echo "usage: $0 <agent-dir> <from> <to>"
	exit 1
fi

set -e

for ((i = $2; i < $3; i++ ))
do
	echo "Starting to background: ""$i"
	./make-agent.sh "$1"-"$i" &
#	sleep 0.9
done

