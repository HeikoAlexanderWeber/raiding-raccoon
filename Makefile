ifeq ($(OS),Windows_NT)
	HOST_OS := windows
	PROGRAM := raiding-raccoon.exe
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        HOST_OS := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        HOST_OS := darwin
    endif
	PROGRAM := raiding-raccoon
endif

ifndef VERSION
	VERSION := $(shell sed -n 1p ./version)
endif

.PHONY: install format build test cover run

install:
	go mod download

format:
	gofmt -s -w .

build:
	rm -rf bin && mkdir bin
	cd src/cmd && go build -o ../../bin/$(PROGRAM) main.go

test:
	go test -v ./test/...

cover:
	go test -v -coverpkg=./src/... -coverprofile ./test/coverage.out ./src/...
	go tool cover -html ./test/coverage.out

run:
	go run src/cmd/main.go --start https://cassini.de
