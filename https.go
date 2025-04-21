package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func main() {
	certPEM, err := os.ReadFile("cert.pem")
	if err != nil {
    		log.Fatal(err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(certPEM)
	if !ok {
    		log.Fatal("failed to parse root certificate")
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
            			RootCAs: roots,
			},
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		panic(err)
	}
	log.Default().Printf("Request status: %d\n", resp.StatusCode)
}

