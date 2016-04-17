#!/bin/sh
set -x

basedir=`pwd`/gopath/src/github.com/ECSTeam/twitter-resource
build_dir=`pwd`/assets

mkdir ${build_dir} > /dev/null

set -e

export GOPATH=`pwd`/gopath:${basedir}/cmd/out/Godeps/_workspace

origbase=`pwd`
cd ${basedir}
go build -o ${build_dir}/check ./cmd/check
go build -o ${build_dir}/in ./cmd/in
go build -o ${build_dir}/out ./cmd/out
cd ${origbase}
