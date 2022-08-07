package service

import (
	"context"
	"time"

	"index-price/domain"
)

func NewGenBarService(
	ctx context.Context,
	streams []domain.PriceStreamSubscriber,
	barTypes []domain.BarType,
) *genBarService {
	return &genBarService{
		ctx:      ctx,
		streams:  streams,
		barTypes: barTypes,
	}
}

var _ domain.GenerateBarService = (*genBarService)(nil)

type genBarService struct {
	ctx      context.Context
	streams  []domain.PriceStreamSubscriber
	barTypes []domain.BarType

	buffersTree map[domain.Ticker]map[domain.BarType]domain.BarTimeSeriesBuffer
}

func (svc genBarService) GetBarItemStream(ticker domain.Ticker, barType domain.BarType) (chan domain.BarItem, error) {
	//TODO implement me

	fake := make(chan domain.BarItem)
	go func() {
		ticker := time.NewTicker(250 * time.Millisecond)

	RunLoop:
		for {
			select {
			case <-svc.ctx.Done():
				break RunLoop
			case <-ticker.C:
				fake <- domain.BarItem{}
			}
		}

		close(fake)
	}()

	return fake, nil
}
