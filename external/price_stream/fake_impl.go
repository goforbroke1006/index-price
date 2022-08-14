package price_stream

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"

	"index-price/domain"
)

// NewFakeExchangeService creates streams to prices.
// Prices received in random time intervals - every 100-200 milliseconds
func NewFakeExchangeService() *fakePriceExchangeService {
	return &fakePriceExchangeService{}
}

type fakePriceExchangeService struct{}

var _ domain.PriceStreamSubscriber = (*fakePriceExchangeService)(nil)

func (f fakePriceExchangeService) SubscribePriceStream(ticker domain.Ticker) (chan domain.TickerPrice, chan error) {
	var (
		pricesCh = make(chan domain.TickerPrice)
		errorsCh = make(chan error)
	)
	go func() {
		for {
			<-time.After(time.Duration(100+100*rand.Float64()) * time.Millisecond)

			errProbability := rand.Int63()
			if errProbability%2 == 0 {
				errorsCh <- errors.New("price generation failed")
				continue
			}

			pricesCh <- domain.TickerPrice{
				Ticker: ticker,
				Time:   time.Now(),
				Price:  rand.Float64(),
			}
		}
	}()

	return pricesCh, errorsCh
}
