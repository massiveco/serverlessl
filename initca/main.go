package initca

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var cfg config.Signing

//Request signing
type Request struct {
	CertificateRequest []byte `json:"certificateRequest"`
	Profile            string `json:"profile"`
}

//Response from lambda
type Response struct {
	Certificate []byte `json:"certificate"`
}

//Signer class for signing a cert
type Signer struct {
	S3     *s3.S3
	Config *config.Signing
	Bucket string
	Prefix string
}

//SignerConfig config the signer
type SignerConfig struct {
}

//New Signer
func New() Signer {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return Signer{
		S3:     s3.New(session),
		Prefix: os.Getenv("serverlessl_S3_PREFIX"),
		Bucket: os.Getenv("serverlessl_S3_BUCKET"),
		Config: &config.Signing{
			Default: &config.SigningProfile{
				Expiry:       helpers.OneYear,
				CAConstraint: config.CAConstraint{IsCA: false},
				Usage:        []string{"signing", "key encipherment", "client auth"},
				ExpiryString: "8760h",
			},
		},
	}
}

//Sign sign a request
func (s Signer) Sign(req signer.SignRequest) ([]byte, error) {
	caPem, caKey, err := s.fetchCA()
	if err != nil {
		return nil, err
	}

	ca, err := helpers.ParseCertificatePEM(caPem)
	if err != nil {
		return nil, err
	}

	key, err := helpers.ParsePrivateKeyPEM(caKey)
	if err != nil {
		return nil, err
	}

	sign, err := local.NewSigner(key, ca, signer.DefaultSigAlgo(key), s.Config)
	if err != nil {
		return nil, err
	}

	cert, err := sign.Sign(req)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func (s Signer) fetchCA() (cert, key []byte, err error) {

	caKeyBuf := bytes.NewBuffer(nil)
	caCertBuf := bytes.NewBuffer(nil)

	err = s.downloadFile("/ca.key", caKeyBuf)
	if err != nil {
		return nil, nil, err
	}
	err = s.downloadFile("/ca.crt", caCertBuf)
	if err != nil {
		return nil, nil, err
	}

	return caCertBuf.Bytes(), caKeyBuf.Bytes(), nil
}

func (s Signer) downloadFile(filename string, buf *bytes.Buffer) error {

	s3Key := fmt.Sprintf("%s%s", s.Prefix, filename)
	s3Object, err := s.S3.GetObject(&s3.GetObjectInput{
		Bucket: &s.Bucket,
		Key:    &s3Key,
	})
	if err != nil {
		return err
	}

	defer s3Object.Body.Close()
	if _, err := io.Copy(buf, s3Object.Body); err != nil {
		return err
	}

	return nil
}
