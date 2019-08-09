VERSION := $(shell git describe --abbrev=0 --tags)

.PHONY: build

all: test

packr2:
	packr2

build-all: packr2
	env GOOS=darwin GOARCH=amd64 go build -ldflags '-X vorta/ui.version=${VERSION}' -o dist/cheat-darwin-amd64 cheat.go
	upx dist/cheat-darwin-amd64
	env GOOS=linux GOARCH=amd64 go build -ldflags '-X vorta/ui.version=${VERSION}' -o dist/cheat-linux-amd64 cheat.go
	upx dist/cheat-linux-amd64

release: build-all
	hub release create \
		--attach=dist/cheat-linux-amd64 \
		--attach=dist/cheat-darwin-amd64 \
		${VERSION}

build:
	go build -ldflags '-X vorta/ui.version=${VERSION}' -o dist/cheat-go cheat.go
