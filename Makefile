CODE := $(shell find . -name '*.go')
DATE := $(shell date -u +"%Y%m%dT%H%M%SZ")
HASH := $(shell git rev-parse --short HEAD 2>/dev/null)
VTAG := $(shell git tag --points-at HEAD)
VERS := $(shell [ -z $(VTAG) ] && echo 'tip' || echo $(VTAG) )

help: ## Show this help screen
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

###########
# Testing #
###########

##@ Testing

test: ## Run unit tests
	go test -race -cover -coverprofile=coverage.out ./...

integration: ## Run integration tests
	go test -tags integration ./...

###############
# Development #
###############

##@ Development

build: test daemon cli ## (default) Run unit tests and build the daemon and CLI binaries.

#############
# Artifacts #
#############

##@ Artifacts

binaries: daemon cli # Build Grid daemon 'gridd' and Grid control CLI 'gridctl'

daemon:
	go build -ldflags "-X main.date=$(DATE) -X main.vers=$(VERS) -X main.hash=$(HASH)" ./cmd/gridd

cli:
	go build -ldflags "-X main.date=$(DATE) -X main.vers=$(VERS) -X main.hash=$(HASH)" ./cmd/gridctl

