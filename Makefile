
bin/slssl:
	go build -o bin/slssl ./cli

install:
	@go get github.com/aws/aws-sdk-go
	@go get github.com/cloudflare/cfssl/...
	@go get github.com/aws/aws-lambda-go/lambda

test: install
	go test ./...