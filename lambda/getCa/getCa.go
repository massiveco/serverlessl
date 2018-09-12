package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	sslinit "github.com/massiveco/serverlessl/getCa"
	"github.com/massiveco/serverlessl/sign"
	"github.com/massiveco/serverlessl/store"
)

var s3Store store.Store

func init() {
	var err error

	s3Store, err = store.NewS3Store(nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler processes signing requests from the serverlessl CLI
func Handler(request sign.Request) (sign.Response, error) {

	cert, err := sslinit.Generate(s3Store)
	if err != nil {
		return sign.Response{}, err
	}
	resp := sign.Response{
		Certificate: cert,
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
