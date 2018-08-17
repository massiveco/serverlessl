package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	sslinit "github.com/massiveco/serverlessl/init"
	"github.com/massiveco/serverlessl/sign"
	"github.com/massiveco/serverlessl/store"
)

type initResponse struct {
	Certificate string `json:"certificate,omitempty"`
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

	cert, err := sslinit.Generate(s3Store)
	if err != nil {
		return initResponse{}, err
	}
	return initResponse{
		Certificate: string(cert[:]),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
