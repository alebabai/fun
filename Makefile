PACKAGES				:= ./...
GOLANGCI_LINT_VERSION	:= v1.27.0

.PHONY: all
all: install test

.PNONY: codegen
codegen:
	go generate -v $(PACKAGES)

.PNONY: fmt
fmt:
	go fmt $(PACKAGES)

.PHONY: deps
deps:
	go mod tidy -v

.PHONY: prepare
prepare: deps codegen fmt

.PHONY: install
install: prepare
	go install -v $(PACKAGES)

.PHONY: test
test: prepare
	go test -v -race $(PACKAGES)

.PHONY: lint
lint: prepare
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:${GOLANGCI_LINT_VERSION} golangci-lint run -v
