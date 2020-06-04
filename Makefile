#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

export APPNAME="Game 1"

build-mac:
	CGO_ENABLED=1 CC=gcc GOOS=darwin GOARCH=amd64 go build -tags static -ldflags "-s -w" -o ${ROOT}/bin/macos/${APPNAME}
