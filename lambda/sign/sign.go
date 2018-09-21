package main

import (
	"os"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cloudflare/cfssl/signer"
	"github.com/massiveco/serverlessl/sign"
	"github.com/massiveco/serverlessl/store"
)

var slsslSign sign.Signer
var profileOverride = os.Getenv("PROFILE_OVERRIDE")

func init() {

	s3Store, err := store.NewS3Store(nil)
	if err != nil {
		log.Fatal(err)
	}
	slsslSign, err = sign.New(s3Store)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler processes signing requests from the serverlessl CLI
func Handler(request sign.Request) (sign.Response, error) {

	profile := request.Profile

	if profileOverride != "" {
		profile = profileOverride
	}

	cert, err := slsslSign.Sign(signer.SignRequest{
		Request: string(request.CertificateRequest),
		Profile: profile,
	})

	if err != nil {
		return sign.Response{}, err
	}

	return sign.Response{
		Certificate: cert,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
