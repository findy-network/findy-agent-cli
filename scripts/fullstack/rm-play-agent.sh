#!/bin/bash

if [[ -z "$1" ]]; then
	echo 'usage: $0 <agent-name1> <agent-name2> ...'
	exit 1
fi

play_path="play/" ./rm-agent.sh $@

