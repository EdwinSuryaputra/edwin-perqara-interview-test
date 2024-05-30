NAME=edwin-perqara-interview-test
VERSION=0.0.1

# Inject .env file
ifneq (,$(wildcard .env))
    include .env
    export
endif

.PHONY: init generate oapigen build run test test_api

init: generate
	go mod tidy

generate: oapigen generate_mocks

oapigen: 
	rm -rf generated
	@echo "Generating files..."
	mkdir generated || true
	oapi-codegen --config=apiconfig.yml api.yml

build: cmd/main.go 
	@echo "Building..."
	go build -o $@ $<

run: init build
	@./build/main

INTERFACES_GO_FILES := $(shell find services repository -name "interfaces.go")
INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:%.go=%.mock.gen.go)
generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

test:
	go clean -testcache
	go test -short -coverprofile coverage.out -short -v ./...

test_api:
	go clean -testcache
	go test ./tests/...