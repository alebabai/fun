PACKAGES				:= ./...
GO_COVER_REPORT			:= coverage.out
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
	go test -v -race -coverprofile=$(GO_COVER_PROFILE) $(PACKAGES)

.PHONY: coverage
coverage: test
	go tool cover -func=$(GO_COVER_PROFILE) -o coverage.txt
	go tool cover -html=$GO_COVER_PROFILEE) -o coverage.html

.PHONY: lint
lint: prepare
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:${GOLANGCI_LINT_VERSION} golangci-lint run -v
