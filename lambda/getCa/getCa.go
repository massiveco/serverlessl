package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	sslinit "github.com/massiveco/serverlessl/getCa"
	"github.com/massiveco/serverlessl/sign"
	"github.com/massiveco/serverlessl/store"
)

// InitResponse response containing the existing or newly created CA
type InitResponse struct {
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
func Handler(request sign.Request) (InitResponse, error) {

	cert, err := sslinit.Generate(s3Store)
	if err != nil {
		return InitResponse{}, err
	}

	return InitResponse{
		Certificate: string(cert[:]),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
