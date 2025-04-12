# Define the Go binary name
BIN_NAME=cidgravity-ticker

# Go-related variables
GO=go
GOLINT=golangci-lint
GOTEST=go test -v
GOBUILD=go build

# Lint the code
.PHONY: lint
lint:
	$(GOLINT) run

# Build the project
.PHONY: build
build:
	$(GOBUILD) -o $(BIN_NAME)

# Run tests
.PHONY: test
test:
	$(GOTEST) ./...

# Clean up built files
.PHONY: clean
clean:
	rm -f $(BIN_NAME)

# Run lint, build, and test in sequence
.PHONY: all
all: lint build test