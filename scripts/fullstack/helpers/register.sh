#!/bin/bash

if [[ -z $1 ]]; then
	echo "usage: $0 <agent-dir>"
	exit 1
fi

set -e

for i in {110..399} 
do
	echo "Starting to background: ""$i"
	./make-agent.sh "$1"-"$i" &
#	sleep 0.9
done

