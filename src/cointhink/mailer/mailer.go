package mailer

import "log"
import "bytes"
import "cointhink/config"
import "cointhink/proto"
import "crypto/tls"

import "gopkg.in/gomail.v2"
import "text/template"

func MailAlgorunStopped(to string, token string) {
	m := gomail.NewMessage()
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Schedule stopped due to credit balance.")
	m.SetBody("text/plain", "")

	deliver(m)
}

func MailCreditDebit(to string, token string) {
	m := gomail.NewMessage()
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Schedule used one cointhink credit")
	m.SetBody("text/plain", "")

	deliver(m)
}

type MailCreditBuyData struct {
}

func MailCreditBuy(to string, data MailCreditBuyData) {
	m := gomail.NewMessage()
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Cointhink Credit Purchase Succesful")
	m.SetBody("text/plain", templateTron("One credit has been added to your account.", data))
	deliver(m)
}

func MailToken(to string, token string) {
	m := gomail.NewMessage()
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Web Login")
	m.SetBody("text/plain", "Use the magic link below to sign in to cointhink.\n"+
		config.C.QueryString("http.base_url")+"/?token="+token)

	deliver(m)
}

func MailNotify(notify *proto.Notify) {
	m := gomail.NewMessage()
	m.SetHeader("To", notify.Recipient)
	m.SetHeader("Subject", notify.Summary)
	m.SetBody("text/plain", notify.Detail)

	deliver(m)
}

func deliver(m *gomail.Message) {
	m.SetHeader("From", config.C.QueryString("email.from"))
	m.SetHeader("Bcc", config.C.QueryString("email.bcc"))

	d := gomail.NewDialer(config.C.QueryString("email.smtp"), 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("mailer.deliver: %s %s", m.GetHeader("To"), err)
	} else {
		log.Printf("-* emailed %s %s", m.GetHeader("To"), m.GetHeader("Subject"))
	}
}

func templateTron(tdata string, data interface{}) string {
	tmpl, err := template.New("test").Parse(tdata)
	if err != nil {
	} else {
		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, data)
		if err != nil {
		} else {
			return tpl.String()
		}
	}
	return ""
}
