PACKAGES				:= ./...
GOLANGCI_LINT_VERSION	:= v1.28.0

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
	go test -v -race -coverprofile=coverage.out $(PACKAGES)

.PHONY: coverage
coverage: test
	go tool cover -func=coverage.out -o coverage.txt
	go tool cover -html=coverage.out -o coverage.html

.PHONY: lint
lint: prepare
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:${GOLANGCI_LINT_VERSION} golangci-lint run -v
