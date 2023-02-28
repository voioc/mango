.PHONY: mod,build,run,cover

cur_path:=$(shell pwd)
app_path:=$(cur_path)
app_bin:=$(cur_path)/bin/myapp
app_cover_bin:=$(cur_path)/bin/app_cover

export GO111MODULE=on
# 阿里云proxy
export GOPROXY=https://goproxy.cn
export GOPRIVATE=codeup.aliyun.com



export GOSUMDB=off
export GO111MODULE=on
export CGO_ENABLED=0

#加入git访问权限
$(shell cp $(cur_path)/.netrc  ~/.netrc)


default: mod

mod:
    @go mod tidy -v
    @go mod download

build:
    @go env
    @go vet $(app_path)
    @go build -o $(app_bin) $(app_path)
    @chmod +x $(app_bin)
    @echo build finish

run: build
ifeq ($(shell uname), Darwin)
    RUNNING_ENV=dev DEBUG=1 $(app_bin)
else
    RUNNING_ENV=test $(app_bin)
endif

run-test:
    RUNNING_ENV=test DEBUG=1 $(app_bin)

cover:
    @go env
    @go vet $(app_path)
    @go test -coverpkg="./..." -c -cover -covermode=atomic $(app_path) -o $(app_cover_bin) -gcflags='all=-N -l'

#使用完毕关闭git访问权限
$(shell rm  ~/.netrc)


