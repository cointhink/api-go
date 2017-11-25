package main

import "log"
import "fmt"
import "os"

import "cointhink/config"
import "cointhink/db"
import "cointhink/mailer"
import "cointhink/model/account"
import "cointhink/model/credit_journal"

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
	if cmd == "account" {
		if len(os.Args) > 2 {
			email := os.Args[2]
			doAccount(email)
		} else {
			log.Printf("$ coinctl credit <email>")
		}
	}
	if cmd == "credit" {
		if len(os.Args) > 2 {
			email := os.Args[2]
			doCredit(email)
		} else {
			log.Printf("$ coinctl credit <email>")
		}
	}
}

func doAccount(email string) {
	log.Printf("looking for %v", email)
	account, err := account.FindByEmail(email)
	if err != nil {
		log.Printf("account find err %+v", err)
	} else {
		log.Printf("Account %v email:%v scheduleCredits: %v",
			account.Id, account.Email, account.ScheduleCredits)
	}
}

func doCredit(email string) {
	log.Printf("looking for %v", email)
	account, err := account.FindByEmail(email)
	if err != nil {
		log.Printf("account find err %+v", err)
	} else {
		log.Printf("Account %v %v", account.Email, account.Id)
		c_err := credit_journal.Credit(&account, "coinctl", 1, 0.0)
		if c_err != nil {
			log.Printf("credit_journal.Credit %+v", c_err)
		} else {
			mailer.MailCreditBuy(account.Email)
		}
	}
}
