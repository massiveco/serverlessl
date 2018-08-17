package init

import (
	"bytes"
	"os"

	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/initca"
	"github.com/massiveco/serverlessl/store"
)

var (
	caCommonName = os.Getenv("SERVERLESSL_CA_NAME")
	caGroup      = os.Getenv("SERVERLESSL_CA_GROUP")
	caCountry    = os.Getenv("SERVERLESSL_CA_COUNTRY")
	caCity       = os.Getenv("SERVERLESSL_CA_CITY")
	caState      = os.Getenv("SERVERLESSL_CA_STATE")
)

// Generate a serverlessl CA
func Generate(store store.Store) ([]byte, error) {

	csrRequest := csr.CertificateRequest{
		CN: caCommonName,
		Names: []csr.Name{csr.Name{
			C:  caCountry,
			L:  caCity,
			O:  caGroup,
			OU: "CA",
			ST: caState,
		}},
	}

	caCertBuf := new(bytes.Buffer)

	err := store.FetchFile("/ca.crt", caCertBuf)

	if err == nil {

		return caCertBuf.Bytes(), nil
	}

	cert, csr, key, err := initca.New(&csrRequest)
	if err != nil {
		return nil, err
	}

	err = store.PutPublicFile("/ca.crt", bytes.NewReader(cert))
	if err != nil {
		return nil, err
	}

	err = store.PutFile("/ca.csr", bytes.NewReader(csr))
	if err != nil {
		return nil, err
	}

	err = store.PutFile("/ca.key", bytes.NewReader(key))
	if err != nil {
		return nil, err
	}

	return cert, nil
}
