package main

import "log"
import "fmt"
import "os"

import "cointhink/config"

import "github.com/ogier/pflag"

func main() {
	var err error
	pflag.Parse()
	config_file := *pflag.String("config", "config.hjson", "config file in (h)json")

	// config
	err = config.C.Read(config_file)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s: %s", config_file, err))
	}
	log.Printf("config loaded %s", config_file)

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		doCmd(cmd)
	} else {
		log.Printf("$ coinctl <cmd>")
	}

}

func doCmd(cmd string) {
	if cmd == "credit" {
		if len(os.Args) > 2 {
			acct := os.Args[2]
			log.Printf("credit to %v", acct)
		} else {
			log.Printf("$ coinctl credit <email>")
		}
	}
}
