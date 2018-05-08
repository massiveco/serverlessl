package cmd

import (
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/massiveco/serverlessl/client"
	"github.com/spf13/cobra"
)

var privateFilename, certificateFilename, csrFilename string

var requestCertificateCmd = &cobra.Command{
	Use:   "requestCertificate",
	Short: "Request a certificate from a serverlessl instance",
	Long:  "Generate a Private key, sign a Certificate Request and send the CSR to a serverlessl instance",
	Run: func(cmd *cobra.Command, args []string) {

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
	},
}

func init() {
	rootCmd.AddCommand(requestCertificateCmd)

	requestCertificateCmd.Flags().StringVarP(&privateFilename, "private-key", "p", "cert.key", "Location to write the private key")
	requestCertificateCmd.Flags().StringVarP(&certificateFilename, "cert", "c", "cert.crt", "Location to write the certificate")
	requestCertificateCmd.Flags().StringVarP(&csrFilename, "cert-signing-request", "r", "cert.csr", "Location to write the certificate signing request")
}
