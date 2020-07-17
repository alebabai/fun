PACKAGES := ./...

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
