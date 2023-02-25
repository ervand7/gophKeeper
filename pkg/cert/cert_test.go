package cert

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/credentials"
)

func TestLoadServerCertificateSuccess(t *testing.T) {
	result := LoadServerCertificate()
	assert.Equal(
		t,
		&tls.Config{
			Certificates: []tls.Certificate{
				getCertificate(),
			},
		},
		result,
	)
}

func TestLoadClientCertificateSuccess(t *testing.T) {
	result := LoadClientCertificate()
	assert.Equal(
		t,
		credentials.NewTLS(
			&tls.Config{
				Certificates:       []tls.Certificate{getCertificate()},
				InsecureSkipVerify: true,
			},
		),
		result,
	)
}
