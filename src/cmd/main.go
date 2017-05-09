package main

import (
	"log"
	"net/http"

	"cointhink"
	"db"

	"github.com/ogier/pflag"
)

func main() {
	var err error
	pflag.Parse()
	config_file := *pflag.String("config", "config.hjson", "config file in (h)json")

	// config
	err = cointhink.C.Read(config_file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("config loaded %s", config_file)

	// db
	db_url := cointhink.C.QueryString("db.url")
	err = db.D.Open(db_url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("db open %s", db_url)

	// net
	listen_address := cointhink.C.QueryString("http.listen_address")
	log.Printf("http listening %s", listen_address)
	http.HandleFunc("/", cointhink.Upgrade)

	// http mainloop
	log.Fatal(http.ListenAndServe(listen_address, nil))
}
