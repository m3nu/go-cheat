VERSION := $(shell git describe --abbrev=0 --tags)

.PHONY: build

all: test

packr2:
	packr2

build-all: packr2
	env GOOS=darwin GOARCH=amd64 go build -ldflags '-X vorta/ui.version=${VERSION} -s -w' -o dist/cheat-darwin-amd64-${VERSION} cheat.go
	upx --brute dist/cheat-darwin-amd64
	env GOOS=linux GOARCH=amd64 go build -ldflags '-X vorta/ui.version=${VERSION} -s -w' -o dist/cheat-linux-amd64-${VERSION} cheat.go
	upx --brute dist/cheat-linux-amd64
	nfpm pkg -f build/package/nfpm.yaml --target dist/cheat-${VERSION}.dep
	nfpm pkg -f build/package/nfpm.yaml --target dist/cheat-${VERSION}.rpm

release: build-all
	hub release create \
		--attach=dist/cheat-linux-amd64 \
		--attach=dist/cheat-darwin-amd64 \
		--attach=dist/dist/cheat-${VERSION}.rpm \
		--attach=dist/dist/cheat-${VERSION}.deb \
		${VERSION}

build:
	go build -ldflags '-X vorta/ui.version=${VERSION} -s -w' -o dist/cheat-go cheat.go
