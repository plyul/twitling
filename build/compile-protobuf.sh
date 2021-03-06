#!/bin/bash
set -e

BUILDDIR=$(mktemp -d -t twitling-build-XXXXXXXXXX)
IMAGE_FILE="${BUILDDIR}"/image.bin

echo 'Linting protocol buffers...'
buf lint

echo 'Building protocol buffer...'
buf build ./proto -o "${IMAGE_FILE}"

echo 'Generating Go code for compiled protocol buffers...'
rm -rf generated
mkdir -p generated
protoc --descriptor_set_in="${IMAGE_FILE}" --go_out="${BUILDDIR}" --go-grpc_out="${BUILDDIR}" $(buf ls-files --input "${IMAGE_FILE}")
cp -r $BUILDDIR/twitling/generated/* ./generated
