package mailer

import "net/smtp"
import "log"
import "config"

func MailToken(to string, token string) {
	log.Print("emailing", token)
	err := smtp.SendMail("localhost:25", nil, config.C.QueryString("email.from"), []string{to}, []byte("msg "+token))
	if err != nil {
		log.Print(err)
	}
}
