.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/create handler/board/create/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/get handler/board/get/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/get_all handler/board/get_all/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/update handler/board/update/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/board/delete handler/board/delete/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
