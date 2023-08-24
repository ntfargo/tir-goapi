.PHONY: all build run clean rustlib

BINARY_NAME = tir-apigo
GO = go
SRC_DIR = ./src
CMD_DIR = $(SRC_DIR)/cmd
GO_FILES = $(wildcard $(CMD_DIR)/*.go)

RUST_REPO = https://github.com/teamcodeyard/tir-engine.git
RUST_LIB_DIR = ./tir-engine

all: rustlib build

rustlib:
ifeq ($(OS),Windows_NT)
	@if not exist "$(RUST_LIB_DIR)" ( \
		echo Cloning Rust tir-engine library... && \
		git clone $(RUST_REPO) $(RUST_LIB_DIR) \
	) else ( \
		echo Rust tir-engine library already exists. \
	)
else
	@if [ ! -d "$(RUST_LIB_DIR)" ]; then \
		echo "Cloning Rust tir-engine library..."; \
		git clone $(RUST_REPO) $(RUST_LIB_DIR); \
	else \
		echo "Rust tir-engine library already exists."; \
	fi
endif
	@echo "Building Rust tir-engine library..."
	@cd $(RUST_LIB_DIR) && cargo build --release

build: $(BINARY_NAME)

$(BINARY_NAME): $(GO_FILES) rustlib
	@echo "Building $(BINARY_NAME)..."
	@cd $(CMD_DIR) && $(GO) build -o ../../$(BINARY_NAME)
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
	@cd $(RUST_LIB_DIR) && cargo clean
	@echo "Cleanup complete"