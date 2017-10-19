package cron

import "time"
import "log"
import "encoding/json"
import "fmt"

import "cointhink/config"
import "cointhink/proto"
import "cointhink/common"
import "cointhink/constants"
import "cointhink/lxd"
import "cointhink/mailer"
import "cointhink/container"
import "cointhink/model/schedule"
import "cointhink/model/account"
import "cointhink/model/algorun"

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
	if time.Minute()%int(config.C.QueryNumber("cron.market_prices_minutes")) == 0 {
		go marketPrices(time)
	}

	if time.Minute()%int(config.C.QueryNumber("cron.schedule_reaper_minutes")) == 0 {
		go scheduleReaper(time)
	}

	if day.YearDay() != time.YearDay() {
		day = time
		CronDay(time)
	}
}

func CronDay(time time.Time) {
	common.RespondAll(&proto.TickTock{Time: time.UTC().Format(constants.ISO8601)})
}

func marketPrices(time time.Time) {
	fmt.Printf("coin price fetch.\n")

	pricePing := &proto.MarketPrices{Prices: []*proto.MarketPrice{}}
	coinNames := []string{"bitcoin", "ethereum"}
	for _, coinName := range coinNames {
		coin, err := coinFetch(coinName)
		if err != nil {
			fmt.Printf("priceFetch error for %s %+v\n", coinName, err)
		} else {
			price := &proto.MarketPrice{
				Exchange:   "simulated",
				Market:     coin.Symbol + "/USD",
				Amount:     coin.PriceUsd,
				Currency:   "USD",
				ReceivedAt: time.UTC().Format(constants.ISO8601)}
			pricePing.Prices = append(pricePing.Prices, price)
		}
	}
	fmt.Printf("coin price pump of %d prices.\n", len(pricePing.Prices))
	common.RespondAll(pricePing)
}

func coinFetch(name string) (CoinMarketCap, error) {
	coin := CoinMarketCap{}
	quote_api := "https://api.coinmarketcap.com/v1/ticker/" + name + "/"
	fmt.Println(quote_api)
	now := time.Now()
	bodyJson, err := common.Httpget(quote_api)
	if err != nil {
		fmt.Printf("price fetch error %+v", err)
		return coin, err
	} else {
		list := []CoinMarketCap{}
		err = json.Unmarshal([]byte(bodyJson), &list)
		delay := time.Now().Sub(now).Nanoseconds() * 1000
		go common.InfluxWrite("marketdata", "exchange", "coinmarketcap", string(delay))
		return list[0], err
	}
}

type CoinMarketCap struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
	PriceEur         string `json:"price_eur"`
	Two4HVolumeEur   string `json:"24h_volume_eur"`
	MarketCapEur     string `json:"market_cap_eur"`
}

func scheduleReaper(time time.Time) {
	expiredSchedules := schedule.RunningExpireds(time)
	log.Printf("** cron.scheduleReaper found %d expired schedules.", len(expiredSchedules))
	for _, _schedule := range expiredSchedules {
		_account, err := account.Find(_schedule.AccountId)
		if err != nil {
		} else {
			log.Printf("Schedule %s expired. Account %s.", _schedule.Id, _account.Id)
			err = schedule.EnableUntilNextMonth(_schedule, &_account)
			if err != nil {
				log.Printf("No credits left. Stopping %s", _schedule.Id)
				schedule.UpdateStatus(_schedule, proto.Schedule_disabled)
				mailer.MailScheduleStopped(_account.Email, _schedule.AlgorithmId)
				boxes, _ := algorun.FindReady(_account.Id, _schedule.Id)
				for _, box := range boxes {
					algorun.UpdateStatus(box, proto.Algorun_deleted)
					stat, _ := lxd.Status(box.Id)
					if stat.ErrorCode != 404 {
						container.Stop(box)
					}
				}
			} else {
				log.Printf("Debiting 1 credit for %s", _schedule.Id)
				mailer.MailCreditDebit(_account.Email, _schedule.AlgorithmId)
			}
		}
	}
}
