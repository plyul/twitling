#!/bin/bash
set -e
echo "#"

export PATH="$PATH:/protoc/bin"

BUILDDIR=$(mktemp -d -t twitling-build-XXXXXXXXXX)
IMAGE_FILE="${BUILDDIR}"/image.bin

echo '> Linting & building with `buf`'
buf lint
buf build -o "${IMAGE_FILE}"

echo '> Compiling with `protoc`'
protoc --descriptor_set_in="${IMAGE_FILE}" --go_out="$BUILDDIR" $(buf ls-files --input "${IMAGE_FILE}")

echo Copying compiled files
chmod -R 666 "$BUILDDIR"/*
ls -la  "$BUILDDIR"
cp -v -r "$BUILDDIR"/internal ./
