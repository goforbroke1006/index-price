package service

import (
	"context"
	"math"
	"time"

	"index-price/domain"
	"index-price/internal/bar"
	"index-price/pkg/clock"
)

func NewIndexPriceService(subscribers []domain.PriceStreamSubscriber) *indexPriceService {
	return &indexPriceService{subscribers: subscribers}
}

var _ domain.GenerateIndexPriceService = (*indexPriceService)(nil)

type indexPriceService struct {
	subscribers []domain.PriceStreamSubscriber
}

func (svc indexPriceService) GetStream(
	ctx context.Context,
	ticker domain.Ticker,
	barType domain.BarType,
) (chan domain.IndexPrice, error) {
	var (
		size     = len(svc.subscribers)
		duration = domain.BarTypeToDuration(barType)
		buffers  = make([]domain.TickPriceSamplesBuffer, size)
	)

	// init subscriptions
	for idx, sub := range svc.subscribers {
		buffers[idx] = bar.NewSamplesBuffer(duration)

		go func(ctx context.Context, idx int, sub domain.PriceStreamSubscriber) {
			tickerCh, errorsCh := sub.SubscribePriceStream(ticker)
		ReadLoop:
			for {
				select {
				case <-ctx.Done():
					break ReadLoop
				case err := <-errorsCh:
					//zap.L().Error("can't extract price", zap.Error(err))
					_ = err
				case tickPrice := <-tickerCh:
					buffers[idx].Add(tickPrice.Time.Unix(), tickPrice.Price)
				}
			}
		}(ctx, idx, sub)
	}

	output := make(chan domain.IndexPrice)

	go func() {
		clockTicker := clock.NewOnClockTicker(duration)

	GenerateLoop:
		for {
			select {
			case <-ctx.Done():
				break GenerateLoop

			case truncatedTime := <-clockTicker.C:
				prevInterval := truncatedTime.Add(-1 * duration).Unix()
				sum := 0.0
				count := 0
				for idx := 0; idx < size; idx++ {
					v := buffers[idx].Get(prevInterval)
					if math.IsNaN(v) {
						continue
					}
					sum += v
					count++
				}
				avg := sum / float64(count)

				tickerPrice := domain.IndexPrice{
					Timestamp: time.Now().Truncate(duration).Unix(),
					Price:     avg,
				}
				output <- tickerPrice
			}
		}
	}()

	return output, nil
}
