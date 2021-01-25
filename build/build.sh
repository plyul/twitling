#!/bin/bash
set -e

BUILDDIR=$(mktemp -d -t twitling-build-XXXXXXXXXX)
IMAGE_FILE="${BUILDDIR}"/image.bin

echo 'Linting protocol buffers...'
buf lint

echo 'Building protocol buffer...'
buf build ./proto -o "${IMAGE_FILE}"

echo 'Generating Go code for compiled protocol buffers...'
mkdir -p generated
protoc --descriptor_set_in="${IMAGE_FILE}" --go_out=generated $(buf ls-files --input "${IMAGE_FILE}")
