#!/bin/bash

cd "$1"
cli agent listen -t --logging "-logtostderr -v=5" &
cd ..

