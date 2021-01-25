.PHONY: help
help:
	@echo Тут будет подсказка

.PHONY: build-docker
docker-builder:
	docker build -t builder -f build/Dockerfile .

.PHONY: build
build: docker-builder
	docker run --rm -v $(PWD):/app builder /build.sh
	chown -R $(id -un):$(id -un) .
