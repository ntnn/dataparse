GO ?= go

GOLINT = github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0

default: check

generate:
	find . \( -name '*_gen.go' -o -name '*_gen_test.go' \) -delete
	$(GO) generate ./...

check:
	$(GO) vet ./...
	$(GO) run $(GOLINT) run ./...

test:
	$(GO) test -v -parallel 8 -cover ./...
