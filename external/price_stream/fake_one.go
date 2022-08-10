package price_stream

import "index-price/domain"

func NewFakeOne() *fakePriceStreamClientOne {
	return &fakePriceStreamClientOne{}
}

type fakePriceStreamClientOne struct{}

var _ domain.PriceStreamSubscriber = (*fakePriceStreamClientOne)(nil)

func (f fakePriceStreamClientOne) SubscribePriceStream(ticker domain.Ticker) (chan domain.TickerPrice, chan error) {
	//TODO implement me
	panic("implement me")
}
