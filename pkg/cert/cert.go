// Package cert for working with certificates which are necessary for secure gRPC.
package cert

import (
	"crypto/tls"
	"log"
	"path/filepath"
	"runtime"

	"google.golang.org/grpc/credentials"
)

func LoadServerCertificate() *tls.Config {
	cert := getCertificate()
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
}

func LoadClientCertificate() credentials.TransportCredentials {
	cert := getCertificate()
	return credentials.NewTLS(
		&tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	)
}

func getCertificate() tls.Certificate {
	certDir := getCertDir()
	cert, err := tls.LoadX509KeyPair(
		certDir+"/localhost.crt", certDir+"/localhost.key",
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cert
}

func getCertDir() string {
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	return filepath.Join(currentDir, "/../../cert")
}
