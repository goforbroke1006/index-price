package domain

import (
	"context"
	"time"
)

type BarType string

const (
	BarType1s  = BarType("1s")
	BarType5s  = BarType("5s")
	BarType15s = BarType("15s")
	BarType1m  = BarType("1m")
)

func BarTypeToDuration(bt BarType) time.Duration {
	switch bt {
	case BarType1s:
		return 1 * time.Second
	case BarType5s:
		return 5 * time.Second
	case BarType15s:
		return 15 * time.Second
	case BarType1m:
		return time.Minute
	}

	return 0
}

type BarTimeSeriesBuffer interface {
	Add(timestamp int64, value float64)
	Get() float64
}

type IndexPrice struct {
	Timestamp int64
	Price     float64
}

type GenerateIndexPriceService interface {
	GetStream(ctx context.Context, ticker Ticker, barType BarType) (chan IndexPrice, error)
}
