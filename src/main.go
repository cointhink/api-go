package main

import (
	"log"
	"net/http"

	"github.com/ogier/pflag"
)

var addr = pflag.String("addr", "localhost:8085", "http listen address")
var db = Db{}
var config = Config{}

func main() {
	var err error

	// config
	config_file := "config.hjson"
	err = config.read(config_file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("config loaded %s", config_file)

	// db
	db_url := config.queryString("db.url")
	err = db.open(db_url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("db open %s", db_url)

	// net
	log.Printf("http listening %s", *addr)
	pflag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", Upgrade)

	// http mainloop
	log.Fatal(http.ListenAndServe(*addr, nil))
}
