VERSION := $(shell cat ./VERSION)
BINARY_NAME := nbview
GO_LDFLAGS="-s -X github.com/beringresearch/main.version=$(VERSION) -extldflags=-static"

ubuntu:
	@echo "Building ..."
	go clean
	go get
	rm -f bin/ubuntu/$(BINARY_NAME)
	@GOOS=linux go build -ldflags=$(GO_LDFLAGS) -o bin/ubuntu/$(BINARY_NAME) *.go
	cp VERSION bin/ubuntu/
	@echo "finished"

linux:
	@echo "Building ..."
	go clean
	go get
	rm -f bin/linux/$(BINARY_NAME)
	@GOOS=linux CGO_ENABLED=0 go build -ldflags=$(GO_LDFLAGS) -o bin/linux/$(BINARY_NAME) *.go
	cp VERSION bin/linux/
	@echo "finished"

darwin:
	@echo "Building ..."
	go clean
	go get
	rm -f bin/darwin/$(BINARY_NAME)
	@GOOS=darwin go build -ldflags=$(GO_LDFLAGS) -o bin/darwin/$(BINARY_NAME) *.go
	cp VERSION bin/darwin
	@echo "finished"