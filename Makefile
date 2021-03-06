.PHONY: build clean tool lint help

all: build

build:
	# @关闭回声
	swag init
	@go build -v .
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

dev:
	# @关闭回声
	swag init
	go build -v .
	./go-gin-example

run:
    # @关闭回声
	@go build -v .
	sudo systemctl restart gin
prod:
    # @关闭回声
	go build
	./matchmaker -config=/root/go/matchmaker/conf/prod.ini

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf go-gin-example
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
