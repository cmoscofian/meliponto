APP_VERSION=0.0.1
DOC_PATH=src/docs
CONFIG_PATH=/usr/local/etc
BUILD_PATH=bin/meliponto

all: fmt lint install
dev: fmt lint vet clean
stage: dev test

fmt:
	@echo "### Formatting the source code ###"
	@go fmt ./...

lint:
	@echo "### Linting the source code ###"
	@golint ./...

vet:
	@echo "### Checking for code issues ###"
	@go vet ./...

test:
	@echo "### Testing the app ###"
	@go test ./...

clean:
	@echo "### Removing binaries ###"
	@rm -rf bin

build: clean
	@echo "### Building the binary ###"
	@go build -o ${BUILD_PATH}\
		-ldflags="\
			-X github.com/cmoscofian/meliponto/src/command.version=${APP_VERSION}\
			-X github.com/cmoscofian/meliponto/src/context.dirname=${CONFIG_PATH}\
			-X github.com/cmoscofian/meliponto/src/util.docs=${DOC_PATH}\
		"

install: vet build
	@echo "### Copy the config file onto the ${CONFIG_PATH} ###"
	@cp meliponto.json ${CONFIG_PATH}/meliponto.json