.DEFAULT_GOAL: help

GO_WORKBENCH_DOCKER_IMAGE = go-workbench:latest

.PHONY: help
help: ## Показать подсказку
	@printf "\033[33m%s:\033[0m\n" 'Доступные команды'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: docker-builder
docker-builder:
	docker build -t "$(GO_WORKBENCH_DOCKER_IMAGE)" -f build/Dockerfile .

.PHONY: gen_proto
gen_proto: docker-builder ## Сгенерировать код по спецификациям Protocol Buffer
	docker run --rm -v $(PWD):/app "$(GO_WORKBENCH_DOCKER_IMAGE)" /app/build/build.sh
	sudo chown -R $(shell id -un):$(shell id -un) .

.PHONY: go_mod
go_mod:
	go mod tidy
	go mod verify

.PHONY: lint
lint: gen_proto ## Линтинг исходного кода
	docker run --rm -v $(shell pwd):/app:ro -w /app "$(GO_WORKBENCH_DOCKER_IMAGE)" golangci-lint -v run ./...
