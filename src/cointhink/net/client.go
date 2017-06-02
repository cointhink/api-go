package net

import "net/http"
import "crypto/tls"
import "time"
import "log"

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

func lxdCall(path string) (*http.Response, error) {
	url := config.C.QueryString("lxd.api_url") + path
	log.Printf("lxd call %s", url)
	return Client().Get(url)
}

func LxdStatus(name string) (*http.Response, error) {
	log.Printf("lxd status for %s", name)
	return lxdCall("/1.0/containers/" + name)
}
