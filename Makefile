include .env
export $(shell sed 's/=.*//' .env)

.PHONY: help build clean deploy deploy-function remove

# Help command to display the available Makefile commands.
help:
	@echo "Available commands:"
	@echo "  - build:           Build the Golang binary for deployment."
	@echo "  - clean:           Clean up generated files and directories."
	@echo "  - deploy:          Build and deploy the entire serverless application."
	@echo "  - deploy-built:    Deploy the serverless application using a pre-built binary."
	@echo "  - deploy-function: Deploy a specific Lambda function by name."
	@echo "  - remove:          Remove all deployed resources of the serverless application."
	@echo "  - help:            Display this help message."

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/create handler/board/create/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/get handler/board/get/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/get_all handler/board/get_all/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/update handler/board/update/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/delete handler/board/delete/main.go

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/note/create handler/note/create/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/note/get_all handler/note/get_all/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/note/update handler/note/update/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/note/delete handler/note/delete/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

deploy-built:
	sls deploy --verbose

deploy-function: clean build
	sls deploy function -f ${name}

remove: clean
	sls remove --verbose
