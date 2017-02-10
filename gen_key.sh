#!/bin/bash
STATIC_KEY=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
echo "Generating Static Key : $STATIC_KEY"
echo "export STATIC_KEY=$STATIC_KEY" > ~/.profile.d/static.sh
