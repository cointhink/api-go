package credit_journal

import "cointhink/proto"
import "cointhink/model/account"
import "cointhink/db"
import "log"

var Columns = "id, account_id, schedule_id, status, stripe_tx, credit_adjustment, total_usd"
var Fields = ":id, :account_id, :schedule_id, :status, :stripe_tx, :credit_adjustment, :total_usd"
var Table = "credit_journals"

func Insert(item *proto.CreditJournal) error {
	item.Id = db.NewId(Table)
	_, err := db.D.Handle.NamedExec("insert into "+Table+" ("+Columns+") "+"values ("+Fields+")", item)
	if err != nil {
		log.Printf(Table+" Create err: %v", err)
		return err
	}
	return nil
}

func Credit(_account *proto.Account, stripeTx string, amount int32, value float32) error {
	initialCreditJournal := &proto.CreditJournal{
		AccountId:        _account.Id,
		CreditAdjustment: amount,
		StripeTx:         stripeTx,
		TotalUsd:         value}
	log.Printf("creditJournal Credit %+v", initialCreditJournal)
	c_err := Insert(initialCreditJournal)
	if c_err != nil {
		return c_err
	} else {
		a_err := account.AdjustScheduleCredits(_account, amount)
		return a_err
	}
}

func Debit(_account *proto.Account, schedule *proto.Schedule, amount int32) error {
	amount = amount * -1 // debit
	initialCreditJournal := &proto.CreditJournal{
		AccountId:        _account.Id,
		CreditAdjustment: amount,
		ScheduleId:       schedule.Id}
	log.Printf("creditJournal Debit %+v", initialCreditJournal)
	c_err := Insert(initialCreditJournal)
	if c_err != nil {
		return c_err
	} else {
		a_err := account.AdjustScheduleCredits(_account, amount)
		return a_err
	}
}
