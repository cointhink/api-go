package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"cointhink/common"
	"cointhink/config"
	"cointhink/container"
	"cointhink/cron"
	"cointhink/db"
	"cointhink/lxd"
	"cointhink/q"

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

	// rpc
	common.RPCq = make(chan q.RpcMsg)
	q.OUTq = make(chan q.RpcOut)
	q.LXDOPq = make(chan q.AccountOperation, 2)

	// rpc calls from httploop
	log.Printf("starting rpc goroutine")
	go func() {
		for {
			msg := <-common.RPCq
			common.Rpc(&msg)
		}
	}()

	// client out msgs
	log.Printf("starting rpc respond goroutine")
	go func() {
		for {
			out := <-q.OUTq
			common.Respond(&out)
		}
	}()

	// watch LXD operations
	log.Printf("starting lxd goroutine")
	go func() {
		for {
			op := <-q.LXDOPq
			lxd.AddOp(&op)
		}
	}()

	// monitor queue sizes
	go func() {
		for {
			rpcSize, outSize, lxdSize := len(common.RPCq), len(q.OUTq), len(q.LXDOPq)
			if rpcSize > 0 || outSize > 0 || lxdSize > 0 {
				log.Printf("Q sizes RPCq:%d OUTq:%d LXDOPq:%d", rpcSize, outSize, lxdSize)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// cron
	go func() {
		cron.Setup()
		cron.DoEvery(1*time.Minute, cron.CronMinute)
	}()

	container.SyncAll()
	log.Printf("SyncAll done")

	// net
	listen_address := config.C.QueryString("http.listen_address")

	log.Printf("starting http %s", listen_address)
	go common.Httploop(listen_address)

	log.Printf("READY.")
	select {}
}
