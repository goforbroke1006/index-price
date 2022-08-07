package service

import (
	"time"

	"index-price/domain"
)

func NewGenBarService(
	streams []domain.PriceStreamSubscriber,
	barTypes []domain.BarType,
) *genBarService {
	return &genBarService{
		streams:  streams,
		barTypes: barTypes,
	}
}

var _ domain.GenerateBarService = (*genBarService)(nil)

type genBarService struct {
	streams  []domain.PriceStreamSubscriber
	barTypes []domain.BarType

	buffersTree map[domain.Ticker]map[domain.BarType]domain.BarTimeSeriesBuffer
}

func (svc genBarService) GetBarItemStream(ticker domain.Ticker, barType domain.BarType) (chan domain.BarItem, error) {
	//TODO implement me

	fake := make(chan domain.BarItem)
	go func() {
		for {
			fake <- domain.BarItem{}
			time.Sleep(250 * time.Millisecond)
		}
	}()
	return fake, nil
}
