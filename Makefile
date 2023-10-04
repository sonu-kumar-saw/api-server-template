# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# App Directory 
APP_DIR = cmd/app
# Name of the binary
BINARY_NAME = ../../build/app

# Build target
all: clean build

build:
	@$(GOBUILD) -C $(APP_DIR) -o $(BINARY_NAME)

clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_NAME)

test:
	@$(GOTEST) ./...

run:
	@$(GOBUILD) -C $(APP_DIR) -o $(BINARY_NAME)
	@./build/app

deps:
	@$(GOGET) -v

.PHONY: all build clean test run deps
