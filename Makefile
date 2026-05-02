GO ?= go

TOOLS_DIR := hack/tools
GOLANGCI_LINT_VER := 2.12.1
GOLANGCI_LINT := $(TOOLS_DIR)/golangci-lint-$(GOLANGCI_LINT_VER)

.PHONY: check
check: lint test ##@ Run all checks

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Run linter
	$(GOLANGCI_LINT) run $(GOLANGCI_LINT_FLAGS) ./...

.PHONY: lint-fix ## Run linter and fix issues
lint-fix: override GOLANGCI_LINT_FLAGS := $(GOLANGCI_LINT_FLAGS) --fix
lint-fix: lint

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

$(GOLANGCI_LINT):
	mkdir -p $(TOOLS_DIR)
	$(GO) tool github.com/ntnn/mindl download -tool golangci-lint -common -out $@ -version $(GOLANGCI_LINT_VER)
