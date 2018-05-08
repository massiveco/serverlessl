package client

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/cloudflare/cfssl/csr"
	"github.com/massiveco/serverlessl/sign"
)

var keyParam = csr.BasicKeyRequest{A: "rsa", S: 2048}

//Client lambda PKI client
type Client struct {
	lambda *lambda.Lambda
	config Config
}

//LambdaConfig Lambda function config
type LambdaConfig struct {
	Name   string
	Region string
}

//Config configures the lambada ca client
type Config struct {
	Lambda LambdaConfig
}

//CertificateDetails details of the request
type CertificateDetails struct {
	CommonName string
	Group      string
	Hosts      []string
}

//New create a new client
func New(cfg Config) Client {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := Client{
		lambda: lambda.New(sess, &aws.Config{Region: &cfg.Lambda.Region}),
		config: cfg,
	}

	return client
}

//RequestCertificate Request a signed certificate
func (c Client) RequestCertificate(details CertificateDetails) (csrPEM []byte, keyPEM []byte, certPEM []byte, err error) {

	var cfg *csr.CAConfig
	csrRequest := csr.CertificateRequest{
		CN: details.CommonName,
		Names: []csr.Name{csr.Name{
			O: details.Group,
		}},
		Hosts:      details.Hosts,
		KeyRequest: &keyParam,
		CA:         cfg,
	}

	g := &csr.Generator{Validator: noopValidator}
	csrPEM, keyPEM, err = g.ProcessRequest(&csrRequest)
	if err != nil {
		return nil, nil, nil, err
	}
	req := sign.Request{
		CertificateRequest: csrPEM,
	}
	resp, err := c.invokeLambda(req)
	if err != nil {
		return nil, nil, nil, err
	}

	return csrPEM, keyPEM, resp.Payload, nil
}

func (c Client) invokeLambda(req sign.Request) (*lambda.InvokeOutput, error) {
	payload, _ := json.Marshal(req)

	resp, err := c.lambda.Invoke(&lambda.InvokeInput{FunctionName: &c.config.Lambda.Name, Payload: payload})
	if err != nil {
		return &lambda.InvokeOutput{}, err
	}

	return resp, nil

}

func noopValidator(req *csr.CertificateRequest) error {
	return nil
}
