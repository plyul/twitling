.DEFAULT_GOAL: help

GO_WORKBENCH_DOCKER_IMAGE = go-workbench:latest

.PHONY: help
help: ## Show help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: docker-builder
docker-builder:
	docker build -t "$(GO_WORKBENCH_DOCKER_IMAGE)" -f build/Dockerfile .

.PHONY: gen_proto
gen_proto: docker-builder ## Generate code based on protobuf specs
	docker run --rm -v $(PWD):/app "$(GO_WORKBENCH_DOCKER_IMAGE)" /app/build/compile-protobuf.sh
	sudo chown -R $(shell id -un):$(shell id -un) .

.PHONY: go_mod
go_mod: ## Tidy and verify go.mod and go.sum files
	go mod tidy
	go mod verify

.PHONY: lint
lint: gen_proto ## Lint source code
	docker run --rm -v $(shell pwd):/app:ro -w /app "$(GO_WORKBENCH_DOCKER_IMAGE)" golangci-lint -v run ./...
