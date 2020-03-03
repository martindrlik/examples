package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"
	"net/http"
)

var (
	addr     = flag.String("addr", ":8095", "listen on the TCP network address addr")
	certFile = flag.String("certFile", "cert.pem", "certificate file for server HTTPS connections")
	keyFile  = flag.String("keyFile", "key.pem", "matching private key file")
)

func main() {
	flag.Parse()
	systemCertPool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(newServer(*addr, systemCertPool).
		ListenAndServeTLS(*certFile, *keyFile))
}

func newServer(addr string, certPool *x509.CertPool) *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
	}
	return &http.Server{
		Addr:      addr,
		Handler:   mux,
		TLSConfig: cfg,
	}
}
