package common

import (
	"cointhink/billing"
	"cointhink/httpclients"
	"cointhink/model/account"
	"cointhink/model/token"
	"cointhink/q"

	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func Httploop(listen_address string) {
	http.HandleFunc("/", Upgrade)
	http.HandleFunc("/stripe", Stripe)
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
		_token, err := token.FindByToken(r.Form["cointhink-token"][0])
		if err != nil {
		} else {
			if _token != nil {
				_account, err := account.Find(_token.AccountId)
				if err != nil {
				} else {
					log.Printf("/stripe FORM %+v", r.Form)
					billing.StripePay(r.Form["stripeToken"][0], _account)
					http.Redirect(w, r, "/", 303)
				}
			}
		}
	}
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	wsocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("websocket upgrade fail:", err)
		return
	}
	log.Printf("*- Open websocket for %s from %s", wsocket.RemoteAddr().String(), r.Header.Get("Origin"))
	wsocket.SetPongHandler(func(m string) error {
		if len(httpclients.Pinglist) > 0 {
			httpclients.Pinglist = httpclients.Pinglist[1:]
		}
		return nil
	})

	_client := httpclients.Httpclient{Socket: wsocket, Out: []*q.RpcOut{}}
	httpclients.Clients[wsocket] = _client
	for {
		_, payload, err := wsocket.ReadMessage()
		if err != nil {
			log.Println("ws_recv err:", err)
			break
		}
		RPCq <- q.RpcMsg{Socket: _client.Socket, Payload: payload}
	}
	log.Printf("wsocket closing %s", wsocket.RemoteAddr())
	wsocket.Close()
	httpclients.Clients.Remove(wsocket)
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
