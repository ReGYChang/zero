#!/usr/bin/env bash

export CGO_ENABLED=0

# Delete the old dir
echo "==> Removing old directory..."
rm -f release/*
rm -rf release/*
mkdir -p release/

# Ensure all remote modules are downloaded and cached
go mod download

echo "==> Building..."
go build -ldflags '-s -w -extldflags "-static -fPIC"' -o release ./cmd/*

# Done!
echo
echo "==> Results:"
ls -hl release/
