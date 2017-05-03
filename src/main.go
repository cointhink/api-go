package main

import (
    "fmt"
    "net/http"
    "log"

    "pb"
    "github.com/ogier/pflag"
)

var addr = pflag.String("addr", "localhost:8085", "http service address")

func main() {
    sfrm := signup_form.SignupForm{Email: "nobody@example.com"}
    fmt.Println(sfrm)

    pflag.Parse()
    log.SetFlags(0)
    http.HandleFunc("/", Upgrade)

    // http mainloop
    log.Fatal(http.ListenAndServe(*addr, nil))
}

