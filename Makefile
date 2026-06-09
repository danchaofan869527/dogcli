.PHONY: build clean install test

# Build configuration
BINARY_NAME=dogcli
VERSION:= $(shell git describe --tags --always --dirty 2>/dev/null || echo "v1.0.0")
LDFLAGS=-ldflags "-X main.version=${VERSION}"

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Platform-specific build targets
PLATFORMS=darwin/amd64 darwin/arm64 linux/amd64 linux/386 linux/arm64

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) .
	@echo "Build complete: ./$(BINARY_NAME)"

clean: ## Clean build artifacts
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

install: ## Install binary to /usr/local/bin
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	sudo cp $(BINARY_NAME) /usr/local/bin/
	sudo chmod +x /usr/local/bin/$(BINARY_NAME)
	@echo "Installed successfully!"

test: ## Run tests
	$(GOTEST) -v ./...

build-all: ## Build for all platforms
	@echo "Building for multiple platforms..."
	@$(foreach PLATFORM,$(PLATFORMS), \
		echo "Building for $(PLATFORM)..."; \
		GOOS=$(word 1,$(subst /, ,$(PLATFORM))) \
		GOARCH=$(word 2,$(subst /, ,$(PLATFORM))) \
		$(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-$(word 1,$(subst /, ,$(PLATFORM)))-$(word 2,$(subst /, ,$(PLATFORM))) .; \
	)
	@echo "Build complete! Check ./dist/"

release: build-all ## Create release artifacts
	@echo "Creating release..."
	cd dist && \
	for binary in $(BINARY_NAME)-*; do \
		shasum -a 256 "$$binary" > "$$binary.sha256sum"; \
		zip "$$binary.zip" "$$binary"; \
	done
	@echo "Release artifacts created in ./dist/"

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
