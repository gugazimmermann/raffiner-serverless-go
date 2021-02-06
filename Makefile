.PHONY: build clean deploy

build:
	go get -u ./...
	env GOOS=linux GOARCH=amd64 go build -o bin/main cmd/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose