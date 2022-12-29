VAR=Hello make
NOW=$(shell date)

NAME := go-pg
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X main.revision=$(REVISION)

.PHONY: task1
task1:
	# ENV=test make task3
	echo "Task1 $(VAR) $${ENV}"
task2:
	echo "Task2 $(NOW)"
task3: task1 task2
	echo "Task3"

## Install deps
.PHONY: deps
deps:
	go get -v -d
## Install dev-deps
.PHONY: dev-deps
dev-deps: deps
	go install github.com/Songmu/make2help/cmd/make2help@latest

# Run test
.PHONE: test
test: deps
	go test ./...

# Lint
.PHONY: lint
lint: dev-deps
	go vet ./...

# build binaries e.g. make bin/myproj
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

# build binary
.PHONY: build
build: bin/myproj

# SHOW help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)