CODE := $(shell find . -name '*.go')
DATE := $(shell date -u +"%Y%m%dT%H%M%SZ")
HASH := $(shell git rev-parse --short HEAD 2>/dev/null)
VTAG := $(shell git tag --points-at HEAD)
VERS := $(shell [ -z $(VTAG) ] && echo 'tip' || echo $(VTAG) )

all: build

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

###############
# Development #
###############

##@ Development

build: gridd gridctl ## (default) build binaries for current OS

check: ## Check for linting errors (requires golangci-lint)
	golangci-lint run -E goimports -E golint -E govet

cluster: ## Set up a local cluster (requires kind,yq)
	./hack/allocate.sh
	./hack/configure.sh

###########
# Testing #
###########

##@ Testing

test: $(CODE) ## Run unit tests
	go test -race -cover -coverprofile=coverage.out ./...

integration: $(CODE) ## Run integration tests
	go test -tags integration ./...

e2e: $(CODE) ## Run e2e tests
	@echo "Not Implemented"

#############
# Artifacts #
#############

##@ Artifacts

gridd: $(CODE) ## Build the 'gridd' Grid daemon
	go build -ldflags "-X main.date=$(DATE) -X main.vers=$(VERS) -X main.hash=$(HASH)" ./cmd/gridd

gridctl: $(CODE) ## Build the 'gridctl' Grid CLI
	go build -ldflags "-X main.date=$(DATE) -X main.vers=$(VERS) -X main.hash=$(HASH)" ./cmd/gridctl

clean: ## Remove all build and test artifacts
	@rm -f ./gridd
	@rm -f ./gridctl
	@rm -f ./coverage.out

