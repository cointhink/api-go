package actions

import (
	"log"
	"time"

	"cointhink/constants"
	"cointhink/mailer"
	"cointhink/model/account"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleCreate(scheduleCreate *proto.ScheduleCreate, accountId string) []gproto.Message {
	var responses []gproto.Message

	_account, err := account.Find(accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleCreateResponse{Ok: false, Message: err.Error()})
	} else {
		_schedule, err := schedule.Find(scheduleCreate.Schedule.Id)
		if err != nil {
			log.Printf("schedule new! %+v %+v", _schedule.Id, err)
			responses = create(responses, &_account, scheduleCreate.Schedule)
		} else {
			log.Printf("schedule update! %+v", _schedule.Id)
			responses = update(responses, &_account, scheduleCreate.Schedule)
		}
	}
	return responses
}

func create(responses []gproto.Message, _account *proto.Account, partialSchedule *proto.Schedule) []gproto.Message {
	if _account.ScheduleCredits > 0 {
		executor := proto.Schedule_container
		//executor := proto.Schedule_lambda
		_schedule := proto.Schedule{AccountId: _account.Id,
			AlgorithmId:  partialSchedule.AlgorithmId,
			Executor:     executor,
			Status:       proto.Schedule_disabled,
			InitialState: partialSchedule.InitialState,
			EnabledUntil: time.Now().UTC().Format(constants.ISO8601)}
		log.Printf("inserting sched state %v", _schedule.Status)
		err := schedule.Insert(&_schedule)
		if err != nil {
			responses = append(responses, &proto.ScheduleCreateResponse{Ok: false, Message: err.Error()})
		} else {
			log.Printf("new schedule %s algorithm %s executor %s", _schedule.Id, _schedule.AlgorithmId, _schedule.Executor)
			c_err := schedule.EnableUntilNextMonth(&_schedule, _account)
			if c_err != nil {
				log.Printf("DoScheduleCreate credit_journal Debit err %+v", c_err)
			} else {
				mailer.MailCreditDebit(_account.Email, _schedule.AlgorithmId)
			}
			// autostart
			responses = append(responses, &proto.ScheduleCreateResponse{Ok: true,
				ScheduleCredits: _account.ScheduleCredits})
			scheduleStart := proto.ScheduleStart{ScheduleId: _schedule.Id}
			responses = append(responses, DoScheduleStart(&scheduleStart, _account.Id)...)
		}
	} else {
		responses = append(responses, &proto.ScheduleCreateResponse{Ok: false,
			Message: "No remaining schedule credits."})
	}
	return responses
}

func update(responses []gproto.Message, _account *proto.Account, item *proto.Schedule) []gproto.Message {
	if item.AccountId == _account.Id {
		schedule.UpdateInitialState(item, item.InitialState)
		responses = append(responses, &proto.ScheduleCreateResponse{Ok: true, ScheduleCredits: _account.ScheduleCredits})
	} else {
		responses = append(responses, &proto.AlgorithmDetailResponse{Ok: false, Message: "No permission"})
	}
	return responses
}
