# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOINST=$(GOCMD) install

MODULE_NAME=github.com/solrac97gr/go-shield

# Binary name
BINARY_NAME=bin/go-shield

#Build for Mac:Intel
build-mac-intel:
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) ./cmd/console

# Build
build-console:
	@$(GOBUILD) -o $(BINARY_NAME) ./cmd/console
	@echo "📦 Build Done"

# Run
run-console: build-console
	@echo "🚀 Running App"
	@./$(BINARY_NAME)