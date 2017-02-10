#!/bin/bash
set -e
go build -o app hackerstalk/server
./app
