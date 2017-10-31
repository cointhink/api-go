package lxd

import "net/http"
import "log"
import "encoding/json"
import "bytes"
import "io/ioutil"
import "reflect"
import "mime/multipart"
import "net/url"

import "cointhink/config"
import "cointhink/q"

func lxdPath(path string) string {
	return config.C.QueryString("lxd.api_url") + path
}

func lxdCall(verb string, path string, bodyParts ...interface{}) (*http.Response, error) {
	url := lxdPath(path)
	log.Printf("lxdCall %s %s", verb, url)
	body := &bytes.Buffer{}
	var mime string
	if len(bodyParts) > 0 {
		var payloadBytes []byte
		if reflect.TypeOf(bodyParts[0]).Kind() == reflect.String {
			payloadBytes = []byte(bodyParts[0].(string))
			mime = "application/octet-stream"
			body.Write(payloadBytes)
		} else if reflect.TypeOf(bodyParts[0]).Name() == "HttpFile" {
			httpFile := bodyParts[0].(HttpFile)
			writer := multipart.NewWriter(body)
			part, _ := writer.CreateFormFile("start", httpFile.Name)
			part.Write([]byte(httpFile.Contents))
			mime = writer.FormDataContentType()
		} else {
			bodywtfs := bodyParts[0].([]interface{})
			if len(bodywtfs) > 0 {
				payloadBytes, _ = json.Marshal(bodywtfs[0])
				mime = "application/json"
				body.Write(payloadBytes)
			}
		}
	}
	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		log.Printf("!lxdCall build err: %v", err)
	}
	if body.Len() > 0 {
		req.Header.Set("Content-Type", mime)
		log.Printf("lxdCall body %s", bytes.NewBuffer(body.Bytes()).String())
	}
	httpResult, err := Client().Do(req)
	if err != nil {
		log.Printf("!lxdCall http error %+v", err)
	} else {
		log.Printf("lxdCall http result %d", httpResult.StatusCode)
	}
	return httpResult, err
}

func lxdCallOperation(verb string, path string, bodyParts ...interface{}) (*q.OperationResponse, error) {
	resp, err := lxdCall(verb, path, bodyParts)
	op := q.OperationResponse{}
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
	log.Printf("lxdCallOperation result: Type:'%s' Status:'%s' Error: '%s'", op.Type, op.Status, op.Error)
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
	//log.Printf("lxd.Status for %s %s", name, op.Metadata.Status)
	return &op, nil
}

//{"name": "test01", "architecture": "x86_64", "profiles": ["default"],
//"source": {"type": "image", "alias": "ubuntuimage"}}
type Lxc struct {
	Name     string    `json:"name"`
	Profiles []string  `json:"profiles"`
	Source   LxcSource `json:"source"`
}

type LxcSource struct {
	Type        string `json:"type"`
	Fingerprint string `json:"fingerprint"`
	Source      string `json:"source"`
	Alias       string `json:"alias"`
}

func Launch(lxc Lxc) *q.OperationResponse {
	op, err := lxdCallOperation("POST", "/1.0/containers", lxc)
	if err != nil {
		log.Printf("lxd Launc err %v", err)
	}
	return op
}

type Starter struct {
	Action   string `json:"action"`
	Timeout  int    `json:"timeout"`
	Force    bool   `json:"force"`
	Stateful bool   `json:"stateful"`
}

func Start(containerId string) *q.OperationResponse {
	starter := Starter{Action: "start"}
	op, err := lxdCallOperation("PUT", "/1.0/containers/"+containerId+"/state", starter)
	if err != nil {
		log.Printf("lxd Start err %v", err)
	}
	return op
}

func Stop(containerId string) *q.OperationResponse {
	starter := Starter{Action: "stop", Force: true}
	op, err := lxdCallOperation("PUT", "/1.0/containers/"+containerId+"/state", starter)
	if err != nil {
		log.Printf("lxd Stop err %v", err)
	}
	return op
}

type Runner struct {
	Command     []string `json:"command"`
	Environment struct {
	} `json:"environment"`
	WaitForWebsocket bool `json:"wait-for-websocket"`
	RecordOutput     bool `json:"record-output"`
	Interactive      bool `json:"interactive"`
	Width            int  `json:"width"`
	Height           int  `json:"height"`
}

func Exec(containerId string, command string) *q.OperationResponse {
	runner := Runner{Command: []string{command}}
	op, err := lxdCallOperation("POST", "/1.0/containers/"+containerId+"/exec", runner)
	if err != nil {
		log.Printf("lxd Exec err %v", err)
	}
	return op
}

func Delete(containerId string) *q.OperationResponse {
	op, err := lxdCallOperation("DELETE", "/1.0/containers/"+containerId)
	if err != nil {
		log.Printf("lxd Delete err %v", err)
	}
	return op
}

type HttpFile struct {
	Name     string
	Contents string
}

func FilePut(containerId, filePath string, contents string) {
	query := url.Values{"path": []string{filePath}}
	result, err := lxdCall("POST", "/1.0/containers/"+containerId+"/files?"+query.Encode(),
		contents)
	if err != nil {
		log.Printf("lxd FilePut err %v", err)
	}
	bodyz, _ := ioutil.ReadAll(result.Body)
	log.Printf("fileput http result body %s", bodyz)
}
