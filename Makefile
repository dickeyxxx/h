.PHONY: build run

default: build

build:
	go build -ldflags "-X main.VERSION $(shell ./version.sh)"

run: build
	./hk $(args)
