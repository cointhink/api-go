package billing

import "log"
import "cointhink/config"
import "cointhink/mailer"
import "cointhink/model/credit_journal"
import "cointhink/model/account"
import "github.com/stripe/stripe-go"
import "github.com/stripe/stripe-go/currency"
import "github.com/stripe/stripe-go/charge"

func StripePay(token string, email string, source string) {
	account, err := account.FindByEmail(email)
	if err != nil {
		log.Printf("Stripe pay email not found %+v", email)
	} else {
		log.Printf("Stripe Pay token %+v email %+v source %+v", token, email, source)
		params := &stripe.ChargeParams{
			Amount:   200,
			Currency: currency.USD,
			Desc:     "Cointhink schedule credit",
		}
		params.SetSource(token)

		stripe.Key = config.C.QueryString("stripe.apikey")
		ch, err := charge.New(params)

		if err != nil {
			log.Printf("stripe charge err %+v", err)
		} else {
			log.Printf("stripe charge %+v PAID: %+v", ch, ch.Paid)
			if ch.Paid {
				c_err := credit_journal.Credit(&account, ch.ID, 1, float32(ch.Amount/100))
				if c_err != nil {
					log.Printf("credit_journal.Credit %+v", c_err)
				} else {
					mailer.MailCreditBuy(account.Email)
				}
			}
		}
	}
}
