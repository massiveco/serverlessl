package main

import (
	"encoding/pem"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/massiveco/serverlessl/client"
)

var privateFilename, certificateFilename, csrFilename string

func main() {
	hostname, _ := os.Hostname()
	flag.StringVar(&privateFilename, "key", "./"+hostname+".key", "Location to save the CSR")
	flag.StringVar(&certificateFilename, "crt", "./"+hostname+".crt", "Location to save the certificate")
	flag.StringVar(&csrFilename, "csr", "./"+hostname+".csr", "Location to save the CSR")

	flag.Parse()

	serverlesslReq := client.New(client.Config{
		Lambda: client.LambdaConfig{
			Name:   "serverlesslSign-default",
			Region: "us-east-2",
		},
	})

	csr, key, crt, err := serverlesslReq.RequestCertificate(client.CertificateDetails{
		CommonName: "system:node:worker-0",
		Group:      "system:nodes",
		Hosts:      []string{"worker-0", "api.kubernetes"},
	})
	if err != nil {
		log.Fatal(err)
	}

	cert := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: crt,
	}
	ioutil.WriteFile(csrFilename, csr, 0600)
	ioutil.WriteFile(privateFilename, key, 0600)
	ioutil.WriteFile(certificateFilename, pem.EncodeToMemory(&cert), 0600)
}
