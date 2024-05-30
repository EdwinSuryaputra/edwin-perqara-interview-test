NAME=edwin-perqara-interview-test
VERSION=0.0.1

# Inject .env file
ifneq (,$(wildcard .env))
    include .env
    export
endif

.PHONY: init oapigen build run

init: oapigen
	go mod tidy

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

