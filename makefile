GO ?= go
PROTOC ?= protoc

PROTO_PATH := $(CURDIR)
API_PATH := $(CURDIR)

BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

DATABASE_DSN="postgresql://zero_test:zero_test@localhost:5432/zero_test?sslmode=disable"
TEST_PACKAGES = ./internal/...

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

test:
	go vet $(TEST_PACKAGES)
	go test -race -cover -coverprofile cover.out $(TEST_PACKAGES)
	go tool cover -func=cover.out | tail -n 1

clean:
	rm -rf bin/ cover.out

# Migrate db up to date
migrate-db-up:
	docker run --rm -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -verbose -path=/migrations/ -database=$(DATABASE_DSN) up

# Revert db migration once a step
migrate-db-down:
	docker run --rm -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -verbose -path=/migrations/ -database=$(DATABASE_DSN) down 1

# Force the current version to the given number. It is used for manually resolving dirty migration flag.
# Ref: https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md#forcing-your-database-version
migrate-db-force-%:
	docker run --rm -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -verbose -path=/migrations/ -database=$(DATABASE_DSN) force $*

# Only used for local auth
init-local-db:
	docker exec cantata-postgres bash -c "psql -U zero_test -d zero_test -f /testdata/init_local_dev.sql"