package mailer

import "log"
import "config"
import "crypto/tls"
import "gopkg.in/gomail.v2"

func MailToken(to string, token string) {
	log.Printf("emailing %s %s", to, token)

	m := gomail.NewMessage()
	m.SetHeader("From", config.C.QueryString("email.from"))
	m.SetHeader("To", to)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Web Login")
	m.SetBody("text/plain", "Use the magic link below to sign in to cointhink.\n"+
		config.C.QueryString("http.base_url")+"/?token="+token)

	d := gomail.NewDialer(config.C.QueryString("email.smtp"), 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		log.Print(err)
	}
}
