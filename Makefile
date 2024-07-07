SHELL := /bin/bash
.PHONY: all build run gotool test clean help show
MK_GOPATH := /mnt/File/GoLand_linux/Go1.22.4/go1.22.4/bin
BINARY = bluebell

# 检查PATH中是否已包含ADD_PATH
ifneq (,$(findstring $(MK_GOPATH),$(PATH)))
    # 如果PATH中已包含ADD_PATH，则不进行任何操作
else
    # 如果PATH中不包含ADD_PATH，则将其添加到PATH中
    export PATH:=$(MK_GOPATH):$(PATH)
    export GOROOT:=/mnt/File/GoLand_linux/Go1.22.4/go1.22.4
    export GOPATH:=/mnt/File/GoLand_linux/Go1.22.4/gomod
    export GO111MODULE:=on
    export CGO_ENABLED:=0
    export GOOS:=linux
    export GOARCH:=amd64
endif

run:
	go run main.go

all: gotool run

build: show
	@echo "编译生成二进制文件: ${GOOS}"
	go build -o ${BINARY} -ldflags="-s -w"  main.go
	@echo "生成二进制文件: ${BINARY}"
	ls -alFh ${BINARY}

show:
	@echo "PATH: ${PATH}"
	@echo "GOROOT: ${GOROOT}"
	@echo "GOPATH: ${GOPATH}"
	@echo "GOOS: ${GOOS}"
	@echo "GOARCH: ${GOARCH}"
	@echo "CGO_ENABLED: ${CGO_ENABLED}"
	@echo "GO111MODULE: ${GO111MODULE}"

gotool:
	go fmt ./...
	go vet ./...

test:
	go test ./...

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "Usage: make [target]"
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make test - 运行 Go 测试"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
	@echo "make show - 显示 Go 环境变量"