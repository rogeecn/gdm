PROJECT="gdm"
VERSION=1.0.0
BUILD=`date +%FT%T%z`

default:
	@echo ${PROJECT}

build:
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build

.PHONY: default build

