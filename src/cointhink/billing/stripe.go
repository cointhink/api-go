package billing

import "log"
import "cointhink/config"
import "github.com/stripe/stripe-go"
import "github.com/stripe/stripe-go/currency"
import "github.com/stripe/stripe-go/charge"

func StripePay(token string, email string, source string) {
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
		log.Printf("stripe charge %+v", ch)
		if ch.Status == "succeeded" {
		}
	}
}
