PACKAGES := ./...

.PHONY: all
all: install test

.PNONY: fmt
fmt:
	go fmt $(PACKAGES)

.PHONY: deps
deps:
	go mod tidy -v

.PHONY: prepare
prepare: deps fmt

.PHONY: install
install: prepare
	go install -v $(PACKAGES)

.PHONY: test
test: prepare
	go test -v -race $(PACKAGES)
