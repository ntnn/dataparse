GO ?= go

.PHONY: check
check: lint test ##@ Run all checks

.PHONY: lint
lint: ##@ Run linter
	$(GO) tool golangci-lint run ./...

.PHONY: test
test: ##@ Run tests
	$(GO) test ./...

.PHONY: clean-generated
clean-generated:
	find . -iname '*.gen.go' -delete

.PHONY: generate
generate: clean-generated ##@ Generate code
	$(GO) generate ./...

.PHONY: help
help:
	@awk 'BEGIN { fs="##@ " } { FS=fs } /:.*##@/ { doc=$$2; FS=":"; $$0=$$0; printf "%s: %s\n", $$1, doc; }' $(MAKEFILE_LIST) | grep -v fs=
