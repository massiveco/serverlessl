package sign

import (
	"bytes"
	"encoding/json"

	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
	"github.com/massiveco/serverlessl/store"
)

const (
	caCertificatePath = "/ca.crt"
	caKeyPath         = "/ca.key"
	caConfigPath      = "/ca-config.json"
)

var cfg config.Signing

// Request signing
type Request struct {
	CertificateRequest []byte `json:"certificateRequest"`
	Profile            string `json:"profile"`
}

// Response from lambda
type Response struct {
	Certificate []byte `json:"certificate"`
}

// Signer class for signing a cert
type Signer struct {
	store  store.Store
	signer *local.Signer
}

// SignerConfig config the signer
type SignerConfig struct {
}

// New Signer
func New(store store.Store) (Signer, error) {
	caPem, caKey, err := fetchCA(store)
	if err != nil {
		return Signer{}, err
	}

	ca, err := helpers.ParseCertificatePEM(caPem)
	if err != nil {
		return Signer{}, err
	}

	key, err := helpers.ParsePrivateKeyPEM(caKey)
	if err != nil {
		return Signer{}, err
	}

	cfg, err := fetchProfiles(store)
	if err != nil {
		return Signer{}, err
	}

	sign, err := local.NewSigner(key, ca, signer.DefaultSigAlgo(key), &cfg)
	if err != nil {
		return Signer{}, err
	}

	return Signer{
		store:  store,
		signer: sign,
	}, nil
}

// Sign sign a request
func (s Signer) Sign(req signer.SignRequest) ([]byte, error) {

	cert, err := s.signer.Sign(req)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func fetchCA(store store.Store) (cert, key []byte, err error) {

	caKeyBuf := new(bytes.Buffer)
	caCertBuf := new(bytes.Buffer)

	err = store.FetchFile(caKeyPath, caKeyBuf)
	if err != nil {
		return nil, nil, err
	}
	err = store.FetchFile(caCertificatePath, caCertBuf)
	if err != nil {
		return nil, nil, err
	}

	return caCertBuf.Bytes(), caKeyBuf.Bytes(), nil
}

func fetchProfiles(store store.Store) (profiles config.Signing, err error) {

	profilesBuf := new(bytes.Buffer)

	err = store.FetchFile(caConfigPath, profilesBuf)
	if err != nil {
		return config.Signing{}, err
	}

	err = json.Unmarshal(profilesBuf.Bytes(), &profiles)
	if err != nil {
		return config.Signing{}, err
	}

	return profiles, nil
}
