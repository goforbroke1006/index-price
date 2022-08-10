package price_stream

import "index-price/domain"

func NewFakeTwo() *fakePriceStreamClientTwo {
	return &fakePriceStreamClientTwo{}
}

type fakePriceStreamClientTwo struct{}

var _ domain.PriceStreamSubscriber = (*fakePriceStreamClientTwo)(nil)

func (f fakePriceStreamClientTwo) SubscribePriceStream(ticker domain.Ticker) (chan domain.TickerPrice, chan error) {
	//TODO implement me
	panic("implement me")
}
