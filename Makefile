VERSION := $(shell git describe --abbrev=0 --tags)

build-all:
	env GOOS=darwin GOARCH=amd64 go build -o dist/cheat-darwin-amd64 cheat.go
	env GOOS=linux GOARCH=amd64 go build -o dist/cheat-linux-amd64 cheat.go
