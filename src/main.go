package main

import (
    "net/http"
    "log"

    "github.com/ogier/pflag"
)

var addr = pflag.String("addr", "localhost:8085", "http service address")

func main() {
    log.Printf("listening %s", *addr)
    pflag.Parse()
    log.SetFlags(0)
    http.HandleFunc("/", Upgrade)

    // http mainloop
    log.Fatal(http.ListenAndServe(*addr, nil))
}

