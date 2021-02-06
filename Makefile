.PHONY: build clean deploy

build:
	go get -u ./...
	env GOOS=linux GOARCH=amd64 go build -o bin/clients cmd/clients/clients.go
	env GOOS=linux GOARCH=amd64 go build -o bin/suppliers cmd/suppliers/suppliers.go
	env GOOS=linux GOARCH=amd64 go build -o bin/products cmd/products/products.go

clean:
	rm -rf ./bin
	rm -rf ./.serverless

deploy: build
	sls deploy --verbose

remove:
	sls remove
	rm -rf ./bin
	rm -rf ./.serverless