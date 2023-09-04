BINARY_NAME := tir-apigo
GO_COMPILER := go
SOURCE_DIR := ./src
COMMAND_DIR := $(SOURCE_DIR)/cmd
GO_SOURCE_FILES := $(wildcard $(COMMAND_DIR)/*.go)
RUST_ZIP_URL := https://cdn.linearfox.com/git/tir-engine.zip?v=4 # Modified tir-engine & tir-engine-grpc for tir-goapi
RUST_LIBRARY_DIR := ./tir-engine
PROTO_DIR := ./tir-engine/proto
PROTO_GO_FILES := $(PROTO_DIR)/tir.pb.go $(PROTO_DIR)/tir_grpc.pb.go

.PHONY: all build run clean download_rustlib extract_rustlib rustlib

all: download_rustlib extract_rustlib build

download_rustlib:
	@echo "Downloading Rust tir-engine library..."
	@curl -o tir-engine.zip $(RUST_ZIP_URL)

extract_rustlib: download_rustlib
ifeq ($(wildcard $(RUST_LIBRARY_DIR)/*),)
	@echo "Extracting Rust tir-engine library..."
ifeq ($(OS),Windows_NT)
	@PowerShell Expand-Archive -Path .\tir-engine.zip -DestinationPath .\$(RUST_LIBRARY_DIR) -Force
else
	@unzip -o -d $(RUST_LIBRARY_DIR) tir-engine.zip
endif
else
	@echo "Rust tir-engine library already extracted."
endif

rustlib: extract_rustlib
	@echo "Building Rust tir-engine library..."
	@cd $(RUST_LIBRARY_DIR) && cargo build --release 

build: rustlib
	@echo "Generating proto go files..."
	@protoc --go_out=$(SOURCE_DIR) --go_opt=paths=source_relative --go-grpc_out=$(SOURCE_DIR) --go-grpc_opt=paths=source_relative $(PROTO_DIR)/tir.proto
	@echo "Proto go files generated successfully."
	@echo "Building $(BINARY_NAME)..."
	@cd $(COMMAND_DIR) && $(GO_COMPILER) build -o ../../$(BINARY_NAME)
	@echo "$(BINARY_NAME) build complete"

run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
ifeq ($(OS),Windows_NT)
	@del /Q /F $(BINARY_NAME)
else
	@rm -f $(BINARY_NAME)
endif
	@cd $(RUST_LIBRARY_DIR) && cargo clean
	@echo "Cleanup complete"