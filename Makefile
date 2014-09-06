.PHONY: build run

default: build

build:
	go build

run: build
	./hk $(args)
