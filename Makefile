APP_VERSION=0.0.0
DOC_PATH=src/docs
CONFIG_PATH=/usr/local/etc
BUILD_PATH=bin/meliponto

.PHONY: fmt
fmt:
	@echo "### Formatting the source code ###"
	@go fmt ./...

.PHONY: lint
lint:
	@echo "### Linting the source code ###"
	@golangci-lint run
	@golint ./...

.PHONY: vet
vet:
	@echo "### Checking for code issues ###"
	@go vet ./...

.PHONY: test
test:
	@echo "### Testing the app ###"
	@go test ./...

.PHONY: clean
clean:
	@echo "### Removing binaries ###"
	@rm -rf bin

.PHONY: pre-commit
pre-commit: clean fmt lint vet test

.PHONY: copy-config
copy-config:
	@echo "### Copy the config file onto the ${CONFIG_PATH} ###"
	@cp meliponto.json ${CONFIG_PATH}/meliponto.json

.PHONY: build
build: clean
	@echo "### Building the binary ###"
	@go build -o ${BUILD_PATH}\
		-ldflags="\
			-X github.com/cmoscofian/meliponto/src/command.version=${APP_VERSION}\
			-X github.com/cmoscofian/meliponto/src/context.dirname=${CONFIG_PATH}\
			-X github.com/cmoscofian/meliponto/src/util.docs=${DOC_PATH}\
		"

.PHONY: install
install: copy-config build
