GO ?= go
PROTOC ?= protoc

PROTO_PATH := $(CURDIR)
API_PATH := $(CURDIR)

BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

.PHONY: build check fmt generate proto

build: generate proto
	@sh -c "'$(CURDIR)/scripts/build.sh'"

check: fmt
	@sh -c "'$(CURDIR)/scripts/staticcheck.sh'"

fmt:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

generate:
	@$(GO) generate ./...

proto:
	@$(PROTOC) --proto_path=$(PROTO_PATH) --go_out=$(API_PATH) --go-grpc_out=$(API_PATH) $(PROTO_PATH)/*.proto
