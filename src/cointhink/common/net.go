package common

import (
	"cointhink/billing"
	"cointhink/config"
	"cointhink/httpclients"
	"cointhink/q"

	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func Httploop(listen_address string) {
	http.HandleFunc("/", Upgrade)
	http.HandleFunc("/stripe", Stripe)
	httpclients.Clients = map[*websocket.Conn]httpclients.Httpclient{}
	http.ListenAndServe(listen_address, nil)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Stripe(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print("/stripe form err", err)
	} else {
		log.Printf("stripe %+v", r.Form)
	}

	billing.StripePay(r.Form["stripeToken"][0],
		r.Form["stripeEmail"][0],
		r.Form["stripeTokenType"][0])
	http.Redirect(w, r, "/", 303)
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s webstocket started", r.Header.Get("Origin"))
	wsocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("websocket upgrade fail:", err)
		return
	}
	wsocket.SetPongHandler(func(m string) error {
		return nil
	})

	_client := httpclients.Httpclient{Socket: wsocket, Out: []*q.RpcOut{}}
	log.Printf("wsocket open %p", wsocket)
	httpclients.Clients[wsocket] = _client
	for {
		_, payload, err := wsocket.ReadMessage()
		if err != nil {
			log.Println("ws_recv err:", err)
			break
		}
		RPCq <- q.RpcMsg{Socket: _client.Socket, Payload: payload}
	}
	log.Printf("wsocket closing %p", wsocket)
	wsocket.Close()
	delete(httpclients.Clients, wsocket)
}

func Httpget(url string) (string, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}
	response, err := client.Get(url)
	if err != nil {
		return "", err
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		return string(body), nil
	}
}

func InfluxWrite(measurement string, tagName string, tagValue string, reading string) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}
	influx_url := config.C.QueryString("influx.url") + "/write?db=" + config.C.QueryString("influx.database")
	data := measurement + "," + tagName + "=" + tagValue + " value=" + reading
	log.Printf("InfluxWrite %s %s\n", influx_url, data)
	response, err := client.Post(influx_url, "", strings.NewReader(data))
	if err != nil {
		log.Printf("Influx post err %v\n", err)
	} else {
		log.Printf("Influx response %s %s\n", response.Proto, response.Status)
	}
}
