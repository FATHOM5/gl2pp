SHELL=bash

include .env

.EXPORT_ALL_VARIABLES:;
.PHONY: default init build clean test start
.DEFAULT_GOAL = default

AWS_PROFILE ?=
AWS ?= aws --profile ${AWS_PROFILE}
PROJECT ?= $(shell basename $$PWD)

default:
	@ mmake help

# install deps
init:
	@ go mod vendor

# build apps
build: clean
	@ go build -o bin/${PROJECT} ./...

clean:
	@ rm -rf bin/${PROJECT}

# run tests
test:
	@ APP_ENV=test go test -v ./...
