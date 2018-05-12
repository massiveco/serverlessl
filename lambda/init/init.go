package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	sslinit "github.com/massiveco/serverlessl/init"
	"github.com/massiveco/serverlessl/sign"
	"github.com/massiveco/serverlessl/store"
)

type initResponse struct {
}

var s3Store store.Store

func init() {
	var err error

	s3Store, err = store.NewS3Store(nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler processes signing requests from the serverlessl CLI
func Handler(request sign.Request) (initResponse, error) {

	err := sslinit.Generate(s3Store)
	if err != nil {
		return initResponse{}, err
	}
	return initResponse{}, nil
}

func main() {
	lambda.Start(Handler)
}
