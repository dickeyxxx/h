.PHONY: build fmt vet test doc dist gzip

binaries = bin/hk-darwin-amd64 bin/hk-linux-386 bin/hk-linux-amd64 bin/hk-windows-386 bin/hk-windows-amd64

default: build

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go fmt ./...

doc:
	godoc -http=:6060 -index

build: vet fmt
	go build -v -o ./bin/hk

dist: build gzip

gzip: $(binaries)
	$(foreach bin,$(binaries), gzip --stdout $(bin) > $(bin).gz;)

bin/hk-darwin-amd64.gz: bin/hk-darwin-amd64
	gzip --stdout bin/hk > bin/hk.gz

bin/hk-darwin-amd64:
	OS=darwin GOARCH=amd64 go build -v -o $@
bin/hk-linux-386:
	GOOS=linux GOARCH=386 go build -v -o $@
bin/hk-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -v -o $@
bin/hk-windows-386:
	GOOS=windows GOARCH=386 go build -v -o $@
bin/hk-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -v -o $@
