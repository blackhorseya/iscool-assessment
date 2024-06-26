.PHONY: all help
all: help

help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean:  ## remove artifacts
	@rm -rf cover.out result.json bin out **/out
	@echo Successfuly removed artifacts

.PHONY: lint
lint: ## run golangci-lint
	@golangci-lint run ./...

.PHONY: test
test: ## test go binary
	@go test -v ./...

.PHONY: coverage
coverage: ## generate coverage report
	@go test -json -coverprofile=cover.out ./... >result.json
