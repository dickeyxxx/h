.PHONY: build fmt vet test doc dist gzip

V = $(`cat VERSION`)
binaries = dist/releases/$V/hk-darwin-amd64 dist/hk-linux-386 dist/hk-linux-amd64 dist/hk-windows-386 dist/hk-windows-amd64

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
	go build -v -o ./dist/hk

dist: build gzip

gzip: $(binaries)
	$(foreach bin,$(binaries), gzip --stdout $(bin) > $(bin).gz;)

dist/hk-darwin-amd64:
	OS=darwin GOARCH=amd64 go build -v -o $@
dist/hk-linux-386:
	GOOS=linux GOARCH=386 go build -v -o $@
dist/hk-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -v -o $@
dist/hk-windows-386:
	GOOS=windows GOARCH=386 go build -v -o $@
dist/hk-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -v -o $@
