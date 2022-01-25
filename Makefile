SHELL := /bin/bash
BASEDIR = $(shell pwd)

APP_NAME=feishu
APP_VERSION=1.1.1
IMAGE_NAME="catchzeng/${APP_NAME}:${APP_VERSION}"
IMAGE_LATEST="catchzeng/${APP_NAME}:latest"

all: mod fmt imports lint test
first:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fmt:
	gofmt -w .
mod:
	go mod tidy
imports:
	goimports -w .
lint:
	golangci-lint run
.PHONY: test
test:
	sh scripts/test.sh
.PHONY: build
build:
	rm -f feishu
	go build -o feishu main.go
build-mac:
	rm -f feishu feishu-darwin-amd64.zip
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o feishu main.go
	zip feishu-darwin-amd64.zip feishu
build-linux:
	rm -f feishu feishu-linux-amd64.zip
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o feishu main.go
	zip feishu-linux-amd64.zip feishu
build-win:
	rm -f feishu.exe feishu-windows-amd64.zip
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o feishu.exe main.go
	zip feishu-windows-amd64.zip feishu.exe
build-win32:
	rm -f feishu.exe feishu-windows-386.zip
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o feishu.exe main.go
	zip feishu-windows-386.zip feishu.exe
build-release:
	make build-mac
	make build-linux
	make build-win
	make build-win32
	rm -f feishu feishu.exe
build-docker:
	sh build/package/build.sh ${IMAGE_NAME}
push-docker:
	docker tag ${IMAGE_NAME} ${IMAGE_LATEST};
	docker push ${IMAGE_NAME};
	docker push ${IMAGE_LATEST};
help:
	@echo "first - first time"
	@echo "fmt - go format"
	@echo "mod - go mod tidy"
	@echo "imports - go imports"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"
	@echo "build - build binary"
	@echo "build-mac - build mac binary"
	@echo "build-linux - build linux amd64 binary"
	@echo "build-win - build win amd64 binary"
	@echo "build-win32 - build win 386 binary"
	@echo "build-docker - build docker image"
	@echo "push-docker - push docker image to docker hub"
