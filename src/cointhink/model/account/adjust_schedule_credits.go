package account

import "cointhink/db"
import "cointhink/proto"
import "log"

func AdjustScheduleCredits(account *proto.Account, adjustment int32) error {
	newCredits := account.ScheduleCredits + adjustment
	_, err := db.D.Handle.Exec("update "+Table+" set schedule_credits = $1 where id = $2",
		newCredits, account.Id)
	if err != nil {
		log.Printf(Table+" DecrementScheduleCredits err: %v", err)
		return err
	} else {
		log.Printf("AdjustScheduleCredits %s adj:%d was:%d now:%d", account.Id, adjustment, account.ScheduleCredits, newCredits)
		account.ScheduleCredits = newCredits
		return nil
	}
}
