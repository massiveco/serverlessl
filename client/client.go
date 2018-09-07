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
	lambdaSvc *lambda.Lambda
	config    Config
}

//LambdaConfig Lambda function config
type LambdaConfig struct {
	Region string
}

//Config configures the lambada ca client
type Config struct {
	Lambda LambdaConfig
	Name   string
}

//CertificateDetails details of the request
type CertificateDetails struct {
	CommonName string
	Group      string
	Hosts      []string
}

//New create a new client
func New(cfg Config) Client {

	sharedSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := Client{
		lambdaSvc: lambda.New(sharedSession, &aws.Config{Region: &cfg.Lambda.Region}),
		config:    cfg,
	}

	return client
}

//FetchCa Request a signed certificate
func (c Client) FetchCa() ([]byte, error) {
	resp, err := c.lambdaSvc.Invoke(&lambda.InvokeInput{FunctionName: aws.String("slsslGetCa-" + c.config.Name)})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

//RequestCertificate Request a signed certificate
func (c Client) RequestCertificate(details CertificateDetails) (csrPEM []byte, keyPEM []byte, certPEM []byte, err error) {

	var cfg *csr.CAConfig
	csrRequest := csr.CertificateRequest{
		CN: details.CommonName,
		Names: []csr.Name{{
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
	req, err := json.Marshal(sign.Request{
		CertificateRequest: csrPEM,
		Profile:            "sandwich",
	})
	if err != nil {
		return nil, nil, nil, err
	}

	resp, err := c.lambdaSvc.Invoke(&lambda.InvokeInput{FunctionName: aws.String("slsslSign-" + c.config.Name), Payload: req})
	if err != nil {
		return nil, nil, nil, err
	}

	return csrPEM, keyPEM, resp.Payload, nil
}

func noopValidator(req *csr.CertificateRequest) error {
	return nil
}
