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
func Generate(store store.Store) error {

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

	cert, _, key, err := initca.New(&csrRequest)
	if err != nil {
		return err
	}

	err = store.PutFile("/ca.crt", bytes.NewReader(cert))
	if err != nil {
		return err
	}

	err = store.PutFile("/ca.key", bytes.NewReader(key))
	if err != nil {
		return err
	}

	return nil
}
