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
	@for file in $(shell find cmd/func -name 'main.go') ; do \
		target=bin/$${file#cmd/func/}; \
		target=$${target%/main.go}; \
		echo "$${target#bin/}"; \
		env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o $$target $$file; \
	done

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

deploy-built:
	sls deploy --verbose

# deploy-function-{{functionName}}
deploy-function-%: clean build
	sls deploy function -f $*
	
remove: clean
	sls remove --verbose
