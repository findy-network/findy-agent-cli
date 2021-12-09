#!/bin/bash

if [[ -z $1 ]]; then
	echo "usage: $0 <agent-dir>"
	exit 1
fi

set -e

for i in {110..199} 
do
	echo "Starting to background: ""$i"
	helpers/kick-listen.sh "$1"-"$i"
#	sleep 0.9
done

