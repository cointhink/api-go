package lxd

import "net/http"
import "log"
import "crypto/tls"
import "time"

import "cointhink/config"

func NewClient() http.Client {

	certFile := config.C.QueryString("lxd.certFile")
	keyFile := config.C.QueryString("lxd.keyFile")

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	transport := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true,
			Certificates: []tls.Certificate{cert}}}
	client := http.Client{
		Timeout:   (time.Second * 10),
		Transport: &transport}

	return client
}

var clientCache *http.Client

func Client() *http.Client {
	if clientCache == nil {
		fixedClient := NewClient()
		clientCache = &fixedClient
	}
	return clientCache
}
