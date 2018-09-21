package init

import (
	"bytes"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/initca"
	"github.com/massiveco/serverlessl/store"
)

var (
	caCommonName = os.Getenv("SLSSL_CA_NAME")
	caGroup      = os.Getenv("SLSSL_CA_GROUP")
	caCountry    = os.Getenv("SLSSL_CA_COUNTRY")
	caCity       = os.Getenv("SLSSL_CA_CITY")
	caState      = os.Getenv("SLSSL_CA_STATE")
)

// Generate a serverlessl CA
func Generate(store store.Store) ([]byte, error) {

	flog := log.WithFields(log.Fields{
		"f": "Generate",
	})

	csrRequest := csr.CertificateRequest{
		CN: caCommonName,
		Names: []csr.Name{{
			C:  caCountry,
			L:  caCity,
			O:  caGroup,
			OU: "CA",
			ST: caState,
		}},
	}

	caCertBuf := new(bytes.Buffer)
	flog.Debug("Attempting to fetch CA from store")
	err := store.FetchFile("/ca.crt", caCertBuf)

	if err == nil {
		flog.Debug("CA Found. Returning.")
		return caCertBuf.Bytes(), nil
	}
	flog.Debug("CA Not found. Generating.")
	flog.WithFields(log.Fields{
		"C":  caCountry,
		"L":  caCity,
		"O":  caGroup,
		"OU": "CA",
		"ST": caState,
	}).Info("Generating new CA")

	cert, _, key, err := initca.New(&csrRequest)
	if err != nil {
		return nil, err
	}

	flog.Debug("Saving CA certificate to store")
	err = store.PutFile("/ca.crt", bytes.NewReader(cert))
	if err != nil {
		return nil, err
	}

	flog.Debug("Saving CA key to store")
	err = store.PutFile("/ca.key", bytes.NewReader(key))
	if err != nil {
		return nil, err
	}

	return cert, nil
}
