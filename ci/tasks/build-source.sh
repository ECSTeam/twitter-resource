#!/bin/sh
basedir=`pwd`/gopath/src/github.com/ECSTeam/twitter-resource
build_dir=`pwd`/assets

mkdir ${build_dir} > /dev/null

set -e
set -x

export GOPATH=`pwd`/gopath:${basedir}/Godeps/_workspace

go build -o ${build_dir}/check ${basedir}/cmd/check
go build -o ${build_dir}/in ${basedir}/cmd/in
go build -o ${build_dir}/out ${basedir}/cmd/out
