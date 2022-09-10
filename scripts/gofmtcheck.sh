#!/usr/bin/env bash

echo "==> Checking that code complies with go fmt requirements..."

GO_FILES=$(find . -name "*.go" -type f | grep -v "/vendor/")

gofmt -s -l -w ${GO_FILES}
