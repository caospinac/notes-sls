.PHONY: build clean deploy

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

remove: clean
	sls remove --verbose

deploy-function: clean build
	sls deploy function -f ${name}
