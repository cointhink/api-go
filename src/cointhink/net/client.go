package net

import "net/http"
import "time"
import "log"

import "cointhink/config"

var Client = http.Client{
	Timeout: time.Second * 10,
}

func lxdCall(path string) (*http.Response, error) {
	url := config.C.QueryString("lxd.api_url") + path
	log.Printf("lxd call %s", url)
	return Client.Get(url)
}

func LxdStatus(name string) (*http.Response, error) {
	return lxdCall("/1.0/containers/" + name)
}
