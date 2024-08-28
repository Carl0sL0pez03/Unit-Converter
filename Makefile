# Variables
BINARY_NAME = unit-converter
BUILD_DIR = bin
SOURCE_DIR = ./cmd

# Build the binary
build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)