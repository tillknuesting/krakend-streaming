# Makefile to clone KrakenD-CE, build it, copy the binary, and show debug info

# Variables
REPO_URL := https://github.com/krakend/krakend-ce
TEMP_DIR := $(shell mktemp -d)
CURRENT_DIR := $(shell pwd)
BINARY_NAME := krakend

.PHONY: build-krakend clean

build-krakend:
	@echo "Cloning repository..."
	git clone $(REPO_URL) $(TEMP_DIR)

	@echo "Building KrakenD..."
	cd $(TEMP_DIR) && make build

	@echo "Copying binary to current directory..."
	cp -f $(TEMP_DIR)/krakend $(CURRENT_DIR)/$(BINARY_NAME)

	@echo "Gathering and printing debug information..."
	@echo "GO version:"
	@go version
	@echo "\nGLIBC version:"
	@ldd --version | head -n 1
	@echo "\nBinary information:"
	@file $(BINARY_NAME)
	@echo "\nLinked libraries:"
	@ldd $(BINARY_NAME)

	@echo "Cleaning up..."
	rm -rf $(TEMP_DIR)

	@echo "Build complete. Binary '$(BINARY_NAME)' is now in the current directory."

	@echo "\nKrakenD version information:"
	@chmod +x $(BINARY_NAME)
	@./$(BINARY_NAME) version

clean:
	rm -f $(BINARY_NAME)