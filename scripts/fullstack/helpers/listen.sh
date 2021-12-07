#!/bin/bash

if [[ -z $1 ]]; then
	echo "usage: $0 <agent-dir>"
	exit 1
fi

set -e

for i in {1..109} 
do
	echo "Starting to background: ""$i"
	./kick-listen.sh "$1"-"$i"
#	sleep 0.9
done

