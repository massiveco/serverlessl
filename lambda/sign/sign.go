package main

import (
	"os"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cloudflare/cfssl/signer"
	cflog "github.com/cloudflare/cfssl/log"
	"github.com/massiveco/serverlessl/sign"
	"github.com/massiveco/serverlessl/store"
)

var slsslSign sign.Signer
var profile = os.Getenv("PROFILE")

func init() {
	cflog.Level = cflog.LevelDebug
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
