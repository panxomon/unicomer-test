BUILDPATH=$(CURDIR)
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
API_NAME=holiday-service
LOWER_API_NAME=$(shell echo $(API_NAME) | tr A-Z a-z)

dir:
	@echo "full path: " $(BUILDPATH)

.PHONY: build
build:
	@echo "Creating binary..."
	@go build -o main ./cmd/main.go
	@echo " binary generated in build/bin/main"

test:
	@echo "Running tests..."
	@go test ./... -short ${PKG_LIST}
	@echo "OK"


tidy:
	@go mod tidy


dep: tidy
	@echo "Installing dependencies..."
# Basic tools
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/go-critic/go-critic/cmd/gocritic@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/swaggo/swag/cmd/swag@latest

	@go mod download
	@echo "OK"


test_coverage:
	@echo "Generating html coverage file..."
	@go test ./... -short -coverprofile=coverage.out ${PKG_LIST} && go tool cover -html=coverage.out
	@rm coverage.out
	@echo "OK"

run:
	@go run cmd/main.go

fmt: tidy
	@echo "Formatting go code...."
	@go fmt ./... ${PKG_LIST}
	@echo "OK"


lint: tidy
	@echo "Checking code style..."
	@staticcheck ./... ${PKG_LIST}
	@go vet ./... ${PKG_LIST}
	@gocritic check ./... ${PKG_LIST}
	@echo "OK"

race: tidy
	@go test ./... -race -short ${PKG_LIST}


.PHONY: proto

mocks:
	@echo "Generating mocks"
	@rm -rf ./tests/mocks/
	@mockery --all --dir internal --output ./tests/mocks --case underscore
	@echo "OK"

.PHONY: docs
docs:
	@echo "Generating docs..."
	@swag fmt
	@swag init --g ./cmd/main.go --markdownFiles ./docs/api --codeExampleFiles ./docs/examples --parseInternal
	@echo "OK"

docker-context:
	@echo "Setting docker context to default again"
	@docker context use default
	@echo "OK"
