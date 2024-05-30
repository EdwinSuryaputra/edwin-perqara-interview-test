NAME=edwin-perqara-interview-test
VERSION=0.0.1

# Inject .env file
ifneq (,$(wildcard .env))
    include .env
    export
endif

.PHONY: build
build:
	@go build -o $(NAME)

.PHONY: run
run: build
	@./$(NAME)

generate: oapigen 

oapigen: 
	rm -rf generated
	@echo "Generating files..."
	mkdir generated || true
	oapi-codegen --config=apiconfig.yml api.yml

