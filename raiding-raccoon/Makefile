ifeq ($(OS),Windows_NT)
	HOST_OS := windows
	PROGRAM := program.exe
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        HOST_OS := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        HOST_OS := darwin
    endif
	PROGRAM := program
endif

ifndef VENDOR
	VENDOR := haw
endif
ifndef NAME
	NAME := raiding-raccoon
endif
ifndef VERSION
	VERSION := latest
endif

.PHONY: install build run format test bench cover docker-build docker-build-debug

install:
	go mod download

build:
	rm -rf bin && mkdir bin
	#mkdir bin/res && cp -r res/certs bin/res
	cd src/cmd && go build -o ../../bin/$(PROGRAM) main.go

format:
	gofmt -s -w .

test:
	go test -v -coverpkg=./src/... -coverprofile ./test/coverage.out ./src/...

bench:
	go test -run=__ -bench=. -cpuprofile test/profile_cpu.out ./...
	go tool pprof -svg test/profile_cpu.out > test/profile_cpu.svg

cover:
	$(MAKE) test
	go tool cover -html ./test/coverage.out

docker-build:
	echo "Version: $(VERSION)"
	docker build -t $(VENDOR)/$(NAME):$(VERSION) -f docker/Dockerfile .

docker-build-debug:
	echo "Version: $(VERSION)"
	docker build -t $(VENDOR)/$(NAME):$(VERSION) -f docker/Dockerfile.debug .
