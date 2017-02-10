#!/bin/bash

# Build time에 unique한 static key를 생성한다.
# 환경변수 STATIC_KEY으로 사용되며, static resource는 build마다 unique한 path로 사용된다.

STATIC_KEY=$(date | md5sum)
echo "Generating Static Key : $STATIC_KEY"
echo "export STATIC_KEY=$STATIC_KEY" > $PWD/.profile.d/static.sh
