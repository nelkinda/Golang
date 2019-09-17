.PHONY: all
all: build

.PHONY: fmt
fmt:
	go fmt

.PHONY: test
test: fmt
	go test

.PHONY: lint
lint:
	golint

.PHONY: build
build: hello

hello: fmt lint test
	go build -o $@

.PHONY: clean
clean::
	go clean
