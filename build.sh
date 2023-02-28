#!/bin/sh
###
 # @Description: Do not edit
 # @Author: Jianxuesong
 # @Date: 2021-05-31 21:53:23
 # @LastEditors: Jianxuesong
 # @LastEditTime: 2021-05-31 21:54:35
 # @FilePath: /Melon/build.sh
### 
#

#当前版本号,每次更新服务时都必须更新版本号
CurrentVersion=$(date +%m.%d.%H%M)

# pwd

export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/path
export GOPROXY=https://goproxy.io
# export GOCACHE=/tmp
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# go env

go mod download
go mod tidy -v

Path=$(cd `dirname $0`; pwd)
# Mod=$(head -n 1 go.mod)
read _ Mod < go.mod # 读取一行，空格分隔，获得mod名字
Build=$Mod
# echo $Build

ServerName=$(echo $Mod | awk -F/ '{print $3}' | tr 'A-Z' 'a-z') # 获取最后/分隔最后的名字 melon
# echo $ServerName

# ServerName=$(basename `pwd` | tr 'A-Z' 'a-z')
BuildTime=$(date "+%Y-%m-%d %H:%M:%S")
GitCommit=$(git rev-parse --short HEAD || echo unsupported)
GoVersion=`go version`
GoVersion=${GoVersion##*version }


go build -ldflags "-X main.AppVersion=$CurrentVersion \
                   -X main.GitCommit=$GitCommit \
                   -X 'main.BuildTime=$BuildTime' \
                   -X 'main.GoVersion=$GoVersion' \
                   -X main.RunEnv=release " \
        -o ./bin/$ServerName ./main.go


if [ $? -eq 0 ]
then
    $Path/bin/$ServerName -V
fi