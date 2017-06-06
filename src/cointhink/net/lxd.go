package net

import "net/http"
import "crypto/tls"
import "time"
import "log"
import "encoding/json"
import "bytes"
import "io/ioutil"

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

func lxdPath(path string) string {
	return config.C.QueryString("lxd.api_url") + path
}

func lxdCall(path string) (*http.Response, error) {
	url := lxdPath(path)
	log.Printf("lxd get %s", url)
	return Client().Get(url)
}

func lxdPost(path string, json []byte) (*http.Response, error) {
	url := lxdPath(path)
	log.Printf("lxd post %s %s", url, json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	return Client().Do(req)
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
	Type        string `json:"type"`
	Fingerprint string `json:"fingerprint"`
}

func LxdLaunch(lxc Lxc) {
	_json, _ := json.Marshal(lxc)
	resp, err := lxdPost("/1.0/containers", _json)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	op := LaunchResponse{}
	err = json.Unmarshal(body, &op)
	log.Printf("launch resp: %v %v", op.Operation, err)
	resp.Body.Close()
}

type LaunchResponse struct {
	Type       string `json:"type"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Operation  string `json:"operation"`
	ErrorCode  int    `json:"error_code"`
	Error      string `json:"error"`
	Metadata   struct {
		ID         string `json:"id"`
		Class      string `json:"class"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
		Status     string `json:"status"`
		StatusCode int    `json:"status_code"`
		Resources  struct {
			Containers []string `json:"containers"`
		} `json:"resources"`
		Metadata  interface{} `json:"metadata"`
		MayCancel bool        `json:"may_cancel"`
		Err       string      `json:"err"`
	} `json:"metadata"`
}
