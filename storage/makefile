PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

.PHONY: ensure-deps
ensure-deps:
	@echo "=> Syncing dependencies with go mod tidy"
	@go mod tidy

.PHONY: start
start:
	@go run cmd/server/main.go

.PHONY: build-database
build-database:
	@echo "MysqlRoot Passowrd (if don't have ignore): "; \
    read PASS; \
    bash createDB.sh create $$PASS

