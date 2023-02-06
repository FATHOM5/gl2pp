SHELL=bash

include .env

.EXPORT_ALL_VARIABLES:;
.PHONY: default init build clean optimized test install uninstall
.DEFAULT_GOAL = default

PROJECT ?= $(shell basename $$PWD)
VERSION ?= $(shell git describe --always --tags)

default:
	@ mmake help

# install deps
init:
	@ go mod vendor

# build app
build: clean
	@ CGO_ENABLED=0 go build -v -a \
		-tags urfave_cli_no_docs \
		-ldflags '-X main.SemVer=${VERSION}' \
		-o bin/${PROJECT} \
		./...

build_all:
	@ $(MAKE) clean
	@ $(MAKE) build_linux_amd
	@ $(MAKE) build_linux_arm
	@ $(MAKE) build_osx_apple
	@ $(MAKE) build_osx_intel
	@ $(MAKE) build_windows


build_osx_intel:
	@ env \
		CGO_ENABLED=0 \
		GOOS=darwin \
		GOARCH=amd64 \
		go build -v -a \
		-tags urfave_cli_no_docs \
		-ldflags '-X main.SemVer=$(shell git describe --always --tags)' \
		-o bin/gl2pp-darwin-amd64 \
		./...

build_osx_apple:
	@ env \
		CGO_ENABLED=0 \
		GOOS=darwin \
		GOARCH=arm64 \
		go build -v -a \
		-tags urfave_cli_no_docs \
		-ldflags '-X main.SemVer=$(shell git describe --always --tags)' \
		-o bin/gl2pp-darwin-arm64 \
		./...

build_windows:
	@ env \
		CGO_ENABLED=0 \
		GOOS=windows \
		GOARCH=amd64 \
		go build -v -a \
		-tags urfave_cli_no_docs \
		-ldflags '-X main.SemVer=$(shell git describe --always --tags)' \
		-o bin/gl2pp-windows-amd64.exe \
		./...

build_linux_amd:
	@ env \
		CGO_ENABLED=0 \
		GOOS=linux \
		GOARCH=amd64 \
		go build -v -a \
		-tags urfave_cli_no_docs \
		-ldflags '-X main.SemVer=$(shell git describe --always --tags)' \
		-o bin/gl2pp-linux-amd64 \
		./...

build_linux_arm:
	@ env \
		CGO_ENABLED=0 \
		GOOS=linux \
		GOARCH=arm64 \
		go build -v -a \
		-tags urfave_cli_no_docs \
		-ldflags '-X main.SemVer=$(shell git describe --always --tags)' \
		-o bin/gl2pp-linux-arm64 \
		./...

# optimize the build
optimized: build
	@ upx bin/*

clean:
	@ rm -rf bin/${PROJECT} bin/${PROJECT}-*

# run tests
test:
	@ APP_ENV=test go test -v -cover -coverpkg=./... -coverprofile=coverage.txt ./...
	@ APP_ENV=test go tool cover -func coverage.txt

# install the app (to your $GOPATH/bin)
install:
	@ $(MAKE) optimized
	@ cp bin/${PROJECT} $GOPATH/bin/${PROJECT}

# uninstall the app (from your $GOPATH/bin)
uninstall:
	@ rm $GOPATH/bin/${PROJECT}

%:
	@ true
