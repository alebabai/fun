# fun

> Functional types for Golang

[![Build](https://img.shields.io/github/workflow/status/alebabai/fun/CI)](https://github.com/alebabai/fun/actions?query=workflow%3ACI)
[![Go version](https://img.shields.io/github/go-mod/go-version/alebabai/fun)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/alebabai/fun)](https://goreportcard.com/report/github.com/alebabai/fun)
[![Coverage](https://img.shields.io/codecov/c/github/alebabai/fun)](https://codecov.io/github/alebabai/fun)
[![Release](https://img.shields.io/github/release/alebabai/fun.svg)](https://github.com/alebabai/fun/releases)
[![GoDoc](https://godoc.org/github.com/alebabai/fun?status.svg)](https://pkg.go.dev/mod/github.com/alebabai/fun)

## The Roadmap

- [ ] Basic functional types
  - [ ] For `interface{}` type  
    - [x] Supplier
    - [x] Consumer
    - [ ] BiConsumer
    - [ ] Runnable
    - [ ] Predicate
    - [ ] Function
    - [ ] BiFunction
    - [ ] UnaryOperator
    - [ ] BinaryOperator
  - [x] Code generation for any type
    - [x] Choose engine
    - [x] Develop templates
    - [x] Develop data structure
- [x] Infrastructure  
  - [x] Makefile
  - [x] CI/CD
    - [x] Test codegen
    - [x] Run lints
    - [x] Build artifacts
    - [x] Run tests
    - [x] Coverage report (codecov)
    - [x] Add badges to the README
- [ ] Documentation
  - [ ] Review codebase docs
  - [x] Deploy to docs hub (https://pkg.go.dev/)
  - [ ] Add examples to the README
- [ ] Complex function types
  - [ ] Optional
