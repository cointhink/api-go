package lxd

import "net/http"
import "log"
import "encoding/json"
import "bytes"
import "io"
import "io/ioutil"

import "cointhink/config"

func lxdPath(path string) string {
	return config.C.QueryString("lxd.api_url") + path
}

func lxdCall(verb string, path string, bodyParts ...interface{}) (*http.Response, error) {
	url := lxdPath(path)
	log.Printf("lxd %s %s %+v", verb, url, bodyParts)
	var body io.Reader
	body = nil
	if len(bodyParts) > 0 {
		log.Printf("using body index 0 %+v", bodyParts[0])
		bodywtfs := bodyParts[0].([]interface{})
		if len(bodywtfs) > 0 {
			json, _ := json.Marshal(bodywtfs[0])
			log.Printf("using body sub-array 0 %d %s", len(json), json)
			body = bytes.NewBuffer(json)
		}
	}
	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		log.Printf("lxdCall error: %v", err)
	}
	if body != nil {
		log.Printf("setting mime json")
		req.Header.Set("Content-Type", "application/json")
	}
	httpResult, err := Client().Do(req)
	log.Printf("lxdCall http result %d", httpResult.StatusCode)
	return httpResult, err
}

func lxdCallOperation(verb string, path string, bodyParts ...interface{}) (*OperationResponse, error) {
	resp, err := lxdCall(verb, path, bodyParts)
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
	op, err := lxdCallOperation("POST", "/1.0/containers", lxc)
	if err != nil {
		log.Printf("lxd Delete %v", err)
	}
	return op
}

func Delete(containerId string) *OperationResponse {
	log.Printf("lxd delete for %s", containerId)
	op, err := lxdCallOperation("DELETE", "/1.0/containers/"+containerId)
	if err != nil {
		log.Printf("lxd Delete %v", err)
	}
	return op
}
