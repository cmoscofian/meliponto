APP_VERSION=0.0.1
DOC_PATH=src/docs
CONFIG_PATH=/usr/local/etc
BUILD_PATH=bin/meliponto

all: fmt lint vet clean build

fmt:
	@echo "Formatting the source code"
	go fmt ./...

lint:
	@echo "Linting the source code"
	golint ./...

vet:
	@echo "Checking for code issues"
	go vet ./...

clean:
	@echo "Removing binaries"
	rm -rf bin

build: clean
	@echo "### Building the client binary ###"

	go build -o ${BUILD_PATH} -ldflags="-X github.com/cmoscofian/meliponto/src/command.version=${APP_VERSION} -X github.com/cmoscofian/meliponto/src/util.docs=${DOC_PATH} -X github.com/cmoscofian/meliponto/src/context.dirname=${CONFIG_PATH}"
