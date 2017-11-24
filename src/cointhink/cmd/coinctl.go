package main

import "log"
import "fmt"
import "os"

import "cointhink/config"
import "cointhink/db"
import "cointhink/model/account"

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

	// db
	db_url := config.C.QueryString("db.url")
	err = db.D.Open(db_url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("db open %s", db_url)

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
			email := os.Args[2]
			addCredit(email)
		} else {
			log.Printf("$ coinctl credit <email>")
		}
	}
}

func addCredit(email string) {
	log.Printf("looking for %v", email)
	account, err := account.FindByEmail(email)
	if err != nil {
		log.Printf("credit err %+v", err)
	} else {
		log.Printf("Account %v %v", account.Email, account.Id)
	}
}
