package initca

import "github.com/cloudflare/cfssl/csr"

// CAConfig For the CN and O in the CA Certificate
type CAConfig struct {
	CommonName string `json:"name"`
	Group      string `json:"group"`
}

// InitCA Configure a serverlessl deployment
func InitCA(cfg CAConfig) error {

	csrRequest := csr.CertificateRequest{
		CN: cfg.CommonName,
		Names: []csr.Name{csr.Name{
			O: cfg.Group,
		}},
	}

	return nil
}
