package main

import (
	"context"
	"fmt"

	"index-price/domain"
	"index-price/external/price_stream"
	"index-price/internal/service"
	"index-price/pkg"
)

func main() {
	epCtx := pkg.NewSignalContext(context.Background())

	subscriptionOne := price_stream.NewFakeOne()
	subscriptionTwo := price_stream.NewFakeTwo()
	subscribers := []domain.PriceStreamSubscriber{subscriptionOne, subscriptionTwo}

	svc := service.NewGenBarService(subscribers)

	barsStream, err := svc.GetBarItemStream(epCtx, domain.BTCUSDTicker, domain.BarType1Minute)
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
