GO ?= go
BIN ?= bin/api

.PHONY: fmt fmt-check vet test test-race build run clean

fmt:
	$(GO) fmt ./...

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "gofmt needed:"; gofmt -l .; exit 1)

vet:
	$(GO) vet ./...

test:
	$(GO) test -count=1 ./...

test-race:
	CGO_ENABLED=1 $(GO) test -race -count=1 ./...

build:
	mkdir -p bin
	CGO_ENABLED=0 $(GO) build -trimpath -ldflags="-s -w" -o $(BIN) ./cmd/api

run:
	$(GO) run ./cmd/api

clean:
	rm -rf bin
