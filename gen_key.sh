#!/bin/bash
STATIC_KEY=$(date | md5sum)
echo "Generating Static Key : $STATIC_KEY"
echo "export STATIC_KEY=$STATIC_KEY" > ~/.profile.d/static.sh
