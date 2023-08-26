BINARY_NAME := tir-apigo
GO_COMPILER := go
SOURCE_DIR := ./src
COMMAND_DIR := $(SOURCE_DIR)/cmd
GO_SOURCE_FILES := $(wildcard $(COMMAND_DIR)/*.go)
RUST_REPO_URL := https://github.com/Ptrskay3/tir-engine-grpc
RUST_LIBRARY_DIR := ./tir-engine
 
.PHONY: all build run clean rustlib update_submodules
 
all: rustlib update_submodules build
 
rustlib:
	@echo "Ensuring Rust tir-engine library..."
ifeq ($(OS),Windows_NT)
	@if not exist "$(RUST_LIBRARY_DIR)" ( \
		echo Cloning Rust tir-engine library... && \
		git clone --recursive $(RUST_REPO_URL) $(RUST_LIBRARY_DIR) \
	) else ( \
		echo Rust tir-engine library already exists. \
	)
else
	@if [ ! -d "$(RUST_LIBRARY_DIR)" ]; then \
		echo "Cloning Rust tir-engine library..."; \
		git clone --recursive $(RUST_REPO_URL) $(RUST_LIBRARY_DIR); \
	else \
		echo "Rust tir-engine library already exists."; \
	fi
endif
	@echo "Building Rust tir-engine library..."
	@cd $(RUST_LIBRARY_DIR) && cargo build --release
 
update_submodules:
	@echo "Updating submodules..."
	@cd $(RUST_LIBRARY_DIR) && git submodule update --init --recursive
 
build: $(BINARY_NAME)

$(BINARY_NAME): $(GO_SOURCE_FILES) rustlib
	@echo "Building $(BINARY_NAME)..."
	@cd $(COMMAND_DIR) && $(GO_COMPILER) build -o ../../$(BINARY_NAME)
	@echo "$(BINARY_NAME) build complete"
 
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)
 
clean:
	@echo "Cleaning up..."
ifeq ($(OS),Windows_NT)
	@del /Q /F $(BINARY_NAME)
else
	@rm -f $(BINARY_NAME)
endif
	@cd $(RUST_LIBRARY_DIR) && cargo clean
	@echo "Cleanup complete"
