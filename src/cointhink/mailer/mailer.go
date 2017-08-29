package mailer

import "log"
import "cointhink/config"
import "cointhink/proto"
import "crypto/tls"
import "gopkg.in/gomail.v2"

func MailToken(to string, token string) {
	log.Printf("emailing %s %s", to, token)

	m := gomail.NewMessage()
	m.SetHeader("To", to)
	//m.SetHeader("Bcc", config.C.QueryString("email.bcc"))
	m.SetHeader("Subject", "Web Login")
	m.SetBody("text/plain", "Use the magic link below to sign in to cointhink.\n"+
		config.C.QueryString("http.base_url")+"/?token="+token)

	deliver(m)
}

func MailTrade(to string, tradeSignal *proto.TradeSignal) {
	m := gomail.NewMessage()
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Trade Signal")
	m.SetBody("text/plain", "Trade Signal!")

	deliver(m)
}

func deliver(m *gomail.Message) {
	m.SetHeader("From", config.C.QueryString("email.from"))

	d := gomail.NewDialer(config.C.QueryString("email.smtp"), 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		log.Print(err)
	}
}
