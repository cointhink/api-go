package main

import (
	"log"
	"net/http"

	"github.com/ogier/pflag"
)

var addr = pflag.String("addr", "localhost:8085", "http listen address")

func main() {
	var err error

	// config
	config := Config{}
	config.read("config.hjson")

	// db
	db := Db{}
	db_url := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	err = db.open(db_url)
	if err != nil {
		log.Fatal(err)
	}

	// net
	log.Printf("listening %s", *addr)
	pflag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", Upgrade)

	// http mainloop
	log.Fatal(http.ListenAndServe(*addr, nil))
}
