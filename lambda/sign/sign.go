package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cloudflare/cfssl/signer"
	"github.com/massiveco/serverlessl/sign"
)

//Handler processes signing requests from the serverlessl CLI
func Handler(request sign.Request) (sign.Response, error) {

	s := sign.New()

	cert, err := s.Sign(signer.SignRequest{
		Request: string(request.CertificateRequest),
		Profile: request.Profile,
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
