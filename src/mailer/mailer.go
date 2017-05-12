package mailer

import "log"
import "config"
import "gopkg.in/gomail.v2"

func MailToken(to string, token string) {
	log.Printf("emailing %s %s", to, token)

	m := gomail.NewMessage()
	m.SetHeader("From", config.C.QueryString("email.from"))
	m.SetHeader("To", to)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Web Login")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer(config.C.QueryString("email.smtp"), 25, "", "")

	err := d.DialAndSend(m)
	if err != nil {
		log.Print(err)
	}
}
