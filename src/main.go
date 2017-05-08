package main

import (
	"log"
	"net/http"

	"github.com/ogier/pflag"
)

var db = Db{}
var config = Config{}

// command line parameters

func main() {
	var err error
	pflag.Parse()
	config_file := *pflag.String("config", "config.hjson", "config file in (h)json")

	// config
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
	listen_address := config.queryString("http.listen_address")
	log.Printf("http listening %s", listen_address)
	http.HandleFunc("/", Upgrade)

	// http mainloop
	log.Fatal(http.ListenAndServe(listen_address, nil))
}
