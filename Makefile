# Name of your output binary
BINARY_NAME := commit_lint

# Command to run Go
GO := go

# Build for Linux, Windows, and macOS
build: build-linux build-windows build-macos

build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 $(GO) build -o build/$(BINARY_NAME)_linux_amd64

build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 $(GO) build -o build/$(BINARY_NAME)_windows_amd64.exe

build-macos:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 $(GO) build -o build/$(BINARY_NAME)_darwin_amd64

.PHONY: build build-linux build-windows build-macos