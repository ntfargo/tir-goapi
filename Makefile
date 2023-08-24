.PHONY: all build run clean rustlib

BINARY_NAME = tir-apigo
GO = go
SRC_DIR = ./src
CMD_DIR = $(SRC_DIR)/cmd
SERVER_FILE = $(CMD_DIR)/server.go

RUST_REPO = https://github.com/teamcodeyard/tir-engine.git
RUST_LIB_DIR = ./tir-engine

all: rustlib build

rustlib:
	@if not exist "$(RUST_LIB_DIR)" ( \
		echo Cloning Rust tir-engine library... && \
		git clone $(RUST_REPO) $(RUST_LIB_DIR) \
	) else ( \
		echo Rust tir-engine library already exists. \
	)
	@echo Building Rust tir-engine library...
	@cd $(RUST_LIB_DIR) && cargo build --release

build: $(BINARY_NAME)

$(BINARY_NAME): $(SERVER_FILE) rustlib
	@echo Building $(BINARY_NAME)...
	$(GO) build -o $(BINARY_NAME) $(SERVER_FILE)
	@echo $(BINARY_NAME) build complete

run: build
	@echo Running $(BINARY_NAME)...
	./$(BINARY_NAME)

clean:
	@echo Cleaning up...
	@del /Q /F $(BINARY_NAME)
	@cd $(RUST_LIB_DIR) && cargo clean
	@echo Cleanup complete
