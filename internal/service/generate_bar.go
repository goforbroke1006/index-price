package service

import (
	"context"
	"strconv"

	"index-price/domain"
	"index-price/internal/bar"
)

func NewGenBarService(subscribers []domain.PriceStreamSubscriber) *genBarService {
	return &genBarService{subscribers: subscribers}
}

var _ domain.GenerateBarService = (*genBarService)(nil)

type genBarService struct {
	subscribers []domain.PriceStreamSubscriber
}

func (svc genBarService) GetBarItemStream(ctx context.Context, ticker domain.Ticker, barType domain.BarType) (chan domain.BarItem, error) {
	buffers := make([]domain.BarTimeSeriesBuffer, len(svc.subscribers))
	updates := make(chan struct{}, len(svc.subscribers))

	// collect new ticks from all external services
	// and keep data in bunch of buffers
	for idx, sub := range svc.subscribers {
		buffers[idx] = bar.NewRoundBarTimeSeriesBuffer(barType)

		go func(sub domain.PriceStreamSubscriber, buf domain.BarTimeSeriesBuffer) {
			stream, errors := sub.SubscribePriceStream(ticker)

		StreamLoop:
			for {
				select {
				case <-ctx.Done():
					break StreamLoop
				case err := <-errors:
					// TODO: handle err
					_ = err
				case tp := <-stream:
					priceFloat, _ := strconv.ParseFloat(tp.Price, 64)
					buf.Add(tp.Time.Unix(), priceFloat)

					// notify than one of buffer has fresh data
					updates <- struct{}{}
				}
			}
		}(sub, buffers[idx])
	}

	output := make(chan domain.BarItem, len(svc.subscribers))

	// TODO: implement me
	go func() {
	UpdatesLoop:
		for {
			select {
			case <-ctx.Done():
				break UpdatesLoop
			case <-updates:
				// TODO: compress bars from all buffers to single
				output <- domain.BarItem{}
			}
		}
	}()

	return output, nil
}
