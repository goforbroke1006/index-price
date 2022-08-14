package domain

import "time"

type Ticker string

const (
	BTCUSDTicker Ticker = "BTC_USD"
)

type TickerPrice struct {
	Ticker Ticker
	Time   time.Time
	Price  float64
}

type PriceStreamSubscriber interface {
	SubscribePriceStream(Ticker) (chan TickerPrice, chan error)
}
