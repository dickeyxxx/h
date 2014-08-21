.PHONY: build vet test run doc

default: build
	go build -v -o ./bin/hk

build: vet

run: build
	./bin/hk

test:
	go test ./...

vet:
	go vet ./...

doc:
	godoc -http=:6060 -index
