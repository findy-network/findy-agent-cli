#!/bin/bash

# NOTE!! Never ever use this if youn want to check the error codes by yourself
# as we want to do here:
#set -e

# debugging lines for ready to play with
#lichen --template="" $1
#lichen --template="{{range .Modules}}{{range .Module.Licenses}}{{.Name | printf \"%s\n\"}}{{end}}{{end}}" $1

lichen -c lichen-cfg.yaml --template="" $1
result=$?

rm "$1"

if [ $result -gt 0 ]
then
	echo "Licenssing error"
	exit 1
fi

