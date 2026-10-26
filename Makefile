SHELL := /bin/sh

BIN := bin/okto
VERSION ?= dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -w \
	-X github.com/tukesh1/okto/internal/version.Version=$(VERSION) \
	-X github.com/tukesh1/okto/internal/version.Commit=$(COMMIT) \
	-X github.com/tukesh1/okto/internal/version.Date=$(DATE)

GO_FILES := $(shell find . -name '*.go' -not -path './.git/*' -not -path './.opencode/*')

.PHONY: build run test test-race vet fmt check clean

build:
	mkdir -p bin
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/okto

run:
	go run ./cmd/okto

test:
	go test ./...

test-race:
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -w $(GO_FILES)

check:
	@unformatted="$$(gofmt -l $(GO_FILES))"; \
	if [ -n "$$unformatted" ]; then \
		echo "gofmt required for:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi
	go vet ./...
	go test -race ./...

clean:
	rm -rf bin
