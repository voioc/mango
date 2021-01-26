#!/bin/sh
#

#当前版本号,每次更新服务时都必须更新版本号
CurrentVersion=$(date +%m.%d.%H%M)

export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/path
export GOPROXY=https://goproxy.io
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

Path=$(cd `dirname $0`; pwd)
Build=$(basename `pwd` | tr 'A-Z' 'a-z')/app/build

ServerName=$(basename `pwd` | tr 'A-Z' 'a-z')
BuildTime=$(date "+%Y-%m-%d %H:%M:%S")
GitCommit=$(git rev-parse --short HEAD || echo unsupported)
GoVersion=`go version`
GoVersion=${GoVersion##*version }


go build -ldflags "-X $Build.Version=$CurrentVersion \
                   -X $Build.GitCommit=$GitCommit \
                   -X '$Build.BuildTime=$BuildTime' \
                   -X '$Build.GoVersion=$GoVersion' \
                   -X $Build.RunEnv=release " \
        -o ./bin/$ServerName ./server.go


if [ $? -eq 0 ]
then
    $Path/bin/$ServerName -v
fi

