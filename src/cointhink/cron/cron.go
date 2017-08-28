package cron

import "time"
import "log"

import "cointhink/proto"
import "cointhink/common"

var (
	day time.Time
)

func Setup() {
	day = time.Now()
	log.Printf("Current year-day is %d", day.YearDay())
}

func DoEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func CronMinute(time time.Time) {
	if day.YearDay() != time.YearDay() {
		day = time
		CronDay(time)
	}
}

func CronDay(time time.Time) {
	common.RespondAll(&proto.TickTock{Time: time.UTC().Format("2006-01-02T15:04:05Z07:00")})
}
