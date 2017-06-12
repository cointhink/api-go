package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"cointhink/common"
	"cointhink/config"
	"cointhink/container"
	"cointhink/db"
	"cointhink/lxd"

	"github.com/ogier/pflag"
)

func main() {
	var err error
	pflag.Parse()
	config_file := *pflag.String("config", "config.hjson", "config file in (h)json")

	// pid
	pid := os.Getpid()
	ioutil.WriteFile("pid", []byte(strconv.Itoa(pid)), 0644)

	// config
	err = config.C.Read(config_file)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s: %s", config_file, err))
	}
	log.Printf("config loaded %s", config_file)

	// db
	db_url := config.C.QueryString("db.url")
	err = db.D.Open(db_url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("db open %s", db_url)

	container.SyncAll()

	// rpc
	common.RPCq = make(chan common.RpcMsg)
	common.OUTq = make(chan common.Httpclient)
	lxd.LXDOPq = make(chan lxd.AccountOperation)

	// net
	listen_address := config.C.QueryString("http.listen_address")
	go common.Httploop(listen_address)

	go func() {
		for {
			msg := <-common.RPCq
			common.Rpc(msg)
		}
	}()

	// client out msgs queued in one thread
	go func() {
		for {
			client := <-common.OUTq
			common.Pump(client)
		}
	}()

	// watch LXD operations
	go func() {
		for {
			op := <-lxd.LXDOPq
			lxd.AddOp(op)
		}
	}()

	select {}
}
