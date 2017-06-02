package net

import "net/http"
import "crypto/tls"
import "time"
import "log"
import "encoding/json"
import "bytes"

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
	log.Printf("lxd get %s", url)
	return Client().Get(url)
}

func lxdPost(path string, json []byte) (*http.Response, error) {
	url := config.C.QueryString("lxd.api_url") + path
	log.Printf("lxd post %s %s", url, json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func LxdStatus(name string) (*http.Response, error) {
	log.Printf("lxd status for %s", name)
	return lxdCall("/1.0/containers/" + name)
}

//{"name": "test01", "architecture": "x86_64", "profiles": ["default"],
//"source": {"type": "image", "alias": "ubuntuimage"}}
type Lxc struct {
	Name   string    `json:"name"`
	Source LxcSource `json:"source"`
}

type LxcSource struct {
	Type  string `json:"type"`
	Alias string `json:"alias"`
}

func LxdLaunch(lxc Lxc) {
	json, _ := json.Marshal(lxc)
	resp, err := lxdPost("/1.0/containers", json)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
