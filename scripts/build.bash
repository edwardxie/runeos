#! /bin/env sh
APP="runeos"
ADDR=":8080"

# export MARTINI_ENV=production

PWD=`pwd`
export RUNEOS_ROOT=${PWD}
export GOPATH=${PWD}:$GOPATH
go build -ldflags="-s" -o bin/$APP ./src
echo "[${APP} is Running]================================"
bin/${APP} ${1} ${2} ${3} ${4} ${5} ${6}