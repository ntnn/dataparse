GO ?= go

GOLINT = github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0

default: check

generate:
	find . -name '*.gen.go' -delete
	$(GO) generate ./...

check:
	$(GO) vet ./...
	$(GO) run $(GOLINT) run ./...

test:
	$(GO) test -v ./...
