PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

.PHONY: all
all: ensure-deps fmt test

.PHONY: ensure-deps
ensure-deps:
	@echo "=> Syncing dependencies with go mod tidy"
	@go mod tidy

.PHONY: fmt
fmt:
	@echo "=> Executing go fmt"
	@go fmt ./...

.PHONY: test
test:
	@echo "=> Running tests"
	@go test ./... -covermode=atomic -coverpkg=./... -count=1 -race

.PHONY: test-cover
test-cover:
	@echo "=> Running tests and generating report"
	@go test ./... -covermode=atomic -coverprofile=./coverage.out -coverpkg=./... -count=1
	@go tool cover -html=./coverage.out

.PHONY: start
start:
	@go run cmd/server/main.go

.PHONY: build-database
build-database:
	@echo "MysqlRoot Passowrd (if don't have ignore): "; \
    read PASS; \
    curl -s https://raw.githubusercontent.com/bootcamp-go/bootcamps-scripts/main/meli_database.sh | bash  -s create $$PASS

.PHONY: rebuild-database-with-password
rebuild-database-with-password:
	@echo "MysqlRoot Passowrd (if don't have ignore): "; \
    curl -s https://raw.githubusercontent.com/bootcamp-go/bootcamps-scripts/main/meli_database.sh | bash  -s rebuild ${p}
