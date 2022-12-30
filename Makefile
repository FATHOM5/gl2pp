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

# optimize the build
optimized: build
	@ upx bin/*

clean:
	@ rm -rf bin/${PROJECT}

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
