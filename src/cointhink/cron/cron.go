package cron

import "time"
import "log"
import "encoding/json"
import "fmt"

import "cointhink/proto"
import "cointhink/common"

var (
	day time.Time
)
var ISO8601 string = "2006-01-02T15:04:05Z07:00"

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
	if time.Minute()%5 == 0 {
		go marketPrices(time)
	}

	if day.YearDay() != time.YearDay() {
		day = time
		CronDay(time)
	}
}

func CronDay(time time.Time) {
	common.RespondAll(&proto.TickTock{Time: time.UTC().Format(ISO8601)})
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
				ReceivedAt: time.UTC().Format(ISO8601)}
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
	bodyJson, err := common.Httpget(quote_api)
	if err != nil {
		fmt.Printf("price fetch error %+v", err)
		return coin, err
	} else {
		list := []CoinMarketCap{}
		err = json.Unmarshal([]byte(bodyJson), &list)
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
