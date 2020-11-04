#!/bin/bash

if [[ $_ == $0 || $1 == "" ]]; then
	echo "usage: source ""$0" "<schema-name>"
	exit 1
fi

cmdl="go run ../../."

# create schema and cred def
export CRED_ID=$(eval "$cmdl" service creddef create --schema-id `eval "$cmdl" service schema create --name="$1" --version=1.0 --attributes=email ` --tag t1)

echo $CRED_ID

unset cmdl

