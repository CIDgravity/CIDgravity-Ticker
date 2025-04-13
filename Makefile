BIN_NAME=cidgravity-ticker
GO=go
GOLINT=golangci-lint
GOTEST=go test -v
GOBUILD=go build

# Lint
.PHONY: lint
lint:
	go vet ./...
	$(GOLINT) run

# Build
.PHONY: build
build:
	$(GOBUILD) -o $(BIN_NAME)

# Run tests
.PHONY: test
test:
	$(GOTEST) ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

# Deps download
dep:
	go mod download

# Run
run: build
	./${BIN_NAME}

# Clean
.PHONY: clean
clean:
	rm -f $(BIN_NAME)

# Build openapi docs
.PHONE: openapi
openapi:
	redocly build-docs openapi.json --output docs/index.html

# All targets
.PHONY: all
all: lint build test