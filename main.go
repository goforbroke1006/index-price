package main

import (
	"context"
	"fmt"

	"index-price/domain"
	"index-price/internal/service"
	"index-price/pkg"
)

func main() {
	epCtx := pkg.NewSignalContext(context.Background())

	stream1 := &fakeStream{}
	stream2 := &fakeStream{}
	subscribers := []domain.PriceStreamSubscriber{stream1, stream2}
	svc := service.NewGenBarService(epCtx, subscribers, nil)
	barsStream, err := svc.GetBarItemStream(domain.BTCUSDTicker, domain.BarType1Minute)
	if err != nil {
		panic(err)
	}

	go func() {
		for bar := range barsStream {
			fmt.Printf("🔽 %f 🔼 %f ⏩ %f ... %f ⏩ \n", bar.Min, bar.Max, bar.Open, bar.Close)
		}
		fmt.Println("stop streaming...")
	}()

	pkg.WaitForShutdown(epCtx)
}

type fakeStream struct{}

var _ domain.PriceStreamSubscriber = (*fakeStream)(nil)

func (f fakeStream) SubscribePriceStream(ticker domain.Ticker) (chan domain.TickerPrice, chan error) {
	//TODO implement me
	panic("implement me")
}
