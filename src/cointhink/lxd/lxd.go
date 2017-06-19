package lxd

import "net/http"
import "log"
import "encoding/json"
import "bytes"
import "io/ioutil"

import "cointhink/config"

func lxdPath(path string) string {
	return config.C.QueryString("lxd.api_url") + path
}

func lxdCall(verb string, path string) (*http.Response, error) {
	url := lxdPath(path)
	log.Printf("lxd %s %s", verb, url)
	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		log.Printf("%v", err)
	}
	return Client().Do(req)
}

func lxdCallOperation(verb string, path string) (*OperationResponse, error) {
	resp, err := lxdCall(verb, path)
	op := OperationResponse{}
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &op)
	if err != nil {
		return nil, err
	}
	log.Printf("lxd operation: %s %s", op.Type, op.Status)
	resp.Body.Close()
	return &op, nil
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

func Status(name string) (*LxcStatus, error) {
	resp, err := lxdCall("GET", "/1.0/containers/"+name)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	op := LxcStatus{}
	err = json.Unmarshal(body, &op)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	log.Printf("lxd.Status for %s %s", name, op.Metadata.Status)
	return &op, nil
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

func Launch(lxc Lxc) *OperationResponse {
	_json, _ := json.Marshal(lxc)
	resp, err := lxdPost("/1.0/containers", _json)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	op := OperationResponse{}
	err = json.Unmarshal(body, &op)
	log.Printf("lxd.Launch resp: %v %v", op.Operation, err)
	resp.Body.Close()
	return &op
}

func Delete(containerId string) *OperationResponse {
	log.Printf("lxd delete for %s", containerId)
	op, err := lxdCallOperation("DELETE", "/1.0/containers/"+containerId)
	if err != nil {
		log.Printf("lxd Delete %v", err)
	}
	return op
}
