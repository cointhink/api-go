package common

import "time"
import "log"

var (
	day time.Time
)

func CronSetup() {
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
	//log.Printf("chime %+v", time)
}

func CronDay(time time.Time) {
}
