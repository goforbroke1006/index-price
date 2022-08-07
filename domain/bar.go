package domain

import "time"

type BarType string

const (
	BarType1Minute = BarType("1m")
	BarType5Minute = BarType("5m")
)

type BarTimeSeriesBuffer interface {
	Add(timestamp int64, value float64)
	Avg() float64
}

func TruncateWithBarType(barType BarType, timestamp int64) int64 {
	var duration time.Duration
	switch barType {
	case BarType1Minute:
		duration = 1 * time.Minute
	case BarType5Minute:
		duration = 5 * time.Minute
	}

	return time.Unix(timestamp, 0).Truncate(duration).Unix()
}

type BarItem struct {
	Max   float64
	Min   float64
	Open  float64
	Close float64
}

type GenerateBarService interface {
	GetBarItemStream(ticker Ticker, barType BarType) (chan BarItem, error)
}
