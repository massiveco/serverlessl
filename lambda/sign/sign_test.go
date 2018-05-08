package main_test

import (
	"io/ioutil"
	"log"
	"testing"

	main "github.com/massiveco/serverlessl/functions/sign"
	"github.com/massiveco/serverlessl/sign"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	dat, err := ioutil.ReadFile("fixtures/cert.csr")
	if err != nil {
		log.Panic(err)
	}
	tests := []struct {
		request sign.Request
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: sign.Request{CertificateRequest: dat},
			expect:  "",
			err:     nil,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Certificate)
	}
}
