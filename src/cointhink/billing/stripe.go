package billing

import "log"
import "cointhink/config"
import "cointhink/mailer"
import "cointhink/proto"
import "cointhink/model/credit_journal"
import "cointhink/model/account"
import "github.com/stripe/stripe-go"
import "github.com/stripe/stripe-go/currency"
import "github.com/stripe/stripe-go/order"
import "github.com/stripe/stripe-go/customer"

func StripePay(token string, _account proto.Account) {
	stripe.Key = config.C.QueryString("stripe.apikey")

	var _customer *stripe.Customer
	if len(_account.Stripe) > 0 {
		_customer_get, err := customer.Get(_account.Stripe, nil)
		if err != nil {
			log.Printf("!stripe customer get err %+v", err)
		} else {
			_customer = _customer_get
		}
	} else {
		customerParams := &stripe.CustomerParams{
			Email: _account.Email,
		}
		customerParams.SetSource(token)
		_customer_new, err := customer.New(customerParams)
		if err != nil {
			log.Printf("!stripe customer new err %+v", err)
		} else {
			log.Printf("!stripe NEW customer id %+v", _customer_new.ID)
			account.UpdateStripe(&_account, _customer_new.ID)
			_customer = _customer_new
		}
	}

	orderParams := &stripe.OrderParams{
		Currency: currency.USD,
		Items: []*stripe.OrderItemParams{
			&stripe.OrderItemParams{
				Type:   "sku",
				Parent: "credit-2usd",
			},
		}}

	_order, err := order.New(orderParams)
	if err != nil {
		log.Printf("order new err %+v", err)
	} else {
		_orderPay, err := order.Pay(_order.ID, &stripe.OrderPayParams{Customer: _customer.ID})
		if err != nil {
			log.Printf("stripe charge err %+v", err)
		} else {
			log.Printf("stripe charge %+v PAID: %+v", _orderPay.ID, _orderPay.Status)
			if _orderPay.Status == stripe.StatusPaid {
				c_err := credit_journal.Credit(&_account, _orderPay.ID, 1, float32(_orderPay.Amount/100))
				if c_err != nil {
					log.Printf("credit_journal.Credit %+v", c_err)
				} else {
					mailer.MailCreditBuy(_account.Email)
				}
			}
		}
	}
}
