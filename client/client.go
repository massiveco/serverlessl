package client

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/cloudflare/cfssl/csr"
	cfssllog "github.com/cloudflare/cfssl/log"
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

type NullLogger struct {
}

func (n NullLogger) Debug(string)   {}
func (n NullLogger) Crit(string)    {}
func (n NullLogger) Info(string)    {}
func (n NullLogger) Warning(string) {}
func (n NullLogger) Err(string)     {}
func (n NullLogger) Emerg(string)   {}

//New create a new client
func New(cfg Config) Client {

	sharedSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := Client{
		lambdaSvc: lambda.New(sharedSession, &aws.Config{Region: &cfg.Lambda.Region}),
		config:    cfg,
	}

	cfssllog.SetLogger(NullLogger{})

	return client
}

//FetchCa Request a signed certificate
func (c Client) FetchCa() ([]byte, error) {
	log.WithFields(log.Fields{"f": "FetchCa"}).Info("Fetching CA Certificate")
	resp, err := c.lambdaSvc.Invoke(&lambda.InvokeInput{FunctionName: aws.String("slsslGetCa-" + c.config.Name)})
	if err != nil {
		return nil, err
	}

	caResp := sign.Response{}

	err = json.Unmarshal(resp.Payload, &caResp)
	if err != nil {
		return nil, err
	}
	
	return caResp.Certificate, nil
}

//RequestCertificate Request a signed certificate
func (c Client) RequestCertificate(details CertificateDetails) (csrPEM []byte, keyPEM []byte, certPEM []byte, err error) {
	flog := log.WithFields(log.Fields{"f": "RequestCertificate"})

	flog.WithFields(log.Fields{
		"CN":    details.CommonName,
		"Names": details.Group,
		"Hosts": details.Hosts,
	}).Info("Requesting new Certificate")

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

	flog.Debugf("Invoking Lambda: %s", *aws.String("slsslSign-" + c.config.Name))
	resp, err := c.lambdaSvc.Invoke(&lambda.InvokeInput{FunctionName: aws.String("slsslSign-" + c.config.Name), Payload: req})
	if err != nil {
		return nil, nil, nil, err
	}

	return csrPEM, keyPEM, resp.Payload, nil
}

func noopValidator(req *csr.CertificateRequest) error {
	return nil
}
