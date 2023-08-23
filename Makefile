.PHONY: all build run clean
 
BINARY_NAME = tir-apigo
GO = go 
SRC_DIR = ./src
CMD_DIR = $(SRC_DIR)/cmd
SERVER_FILE = $(CMD_DIR)/server.go

all: build

build: $(BINARY_NAME)

$(BINARY_NAME): $(SERVER_FILE)
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BINARY_NAME) $(SERVER_FILE)
	@echo "$(BINARY_NAME) build complete"

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	@echo "Cleanup complete"
