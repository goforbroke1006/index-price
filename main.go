package main

import (
	"context"
	"fmt"
	"index-price/pkg/entrypoint"

	"go.uber.org/zap"

	"index-price/domain"
	"index-price/external/price_stream"
	"index-price/internal/service"
)

const (
	barTypeDemo = domain.BarType5s
	//barTypeDemo = domain.BarType1m
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()
	zap.ReplaceGlobals(logger)

	epCtx := entrypoint.NewSignalContext(context.Background())

	var subscribers []domain.PriceStreamSubscriber
	for i := 0; i < 100; i++ {
		subscribers = append(subscribers, price_stream.NewFakeExchangeService())
	}

	svc := service.NewIndexPriceService(subscribers)

	barsStream, err := svc.GetStream(epCtx, domain.BTCUSDTicker, barTypeDemo)
	if err != nil {
		panic(err)
	}

	go func() {
		for bar := range barsStream {
			fmt.Println(bar.Timestamp, bar.Price)
		}
		fmt.Println("stop streaming...")
	}()

	entrypoint.WaitForShutdown(epCtx)
}
