.PHONY: build clean deploy

build:
	go get -u ./...
	env GOOS=linux GOARCH=amd64 go build -o bin/clients cmd/clients.go

clean:
	rm -rf ./bin
	rm -rf ./serverless

deploy: clean build
	sls deploy --verbose

remove: clean
	sls remove
	rm -rf ./bin