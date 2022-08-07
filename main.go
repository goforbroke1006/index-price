package main

import (
	"fmt"

	"index-price/domain"
	"index-price/internal/service"
)

func main() {
	stream1 := &fakeStream{}
	stream2 := &fakeStream{}
	subscribers := []domain.PriceStreamSubscriber{stream1, stream2}
	svc := service.NewGenBarService(subscribers, nil)
	barsStream, err := svc.GetBarItemStream(domain.BTCUSDTicker, domain.BarType1Minute)
	if err != nil {
		panic(err)
	}

	for bar := range barsStream {
		fmt.Printf("⬇ %f ⬆ %f ⏩ %f ... %f ⏩\n", bar.Min, bar.Max, bar.Open, bar.Close)
	}
}

type fakeStream struct{}

var _ domain.PriceStreamSubscriber = (*fakeStream)(nil)

func (f fakeStream) SubscribePriceStream(ticker domain.Ticker) (chan domain.TickerPrice, chan error) {
	//TODO implement me
	panic("implement me")
}
