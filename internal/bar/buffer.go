package bar

import (
	"math"
	"time"

	"index-price/domain"
)

func NewRoundBarTimeSeriesBuffer(interval time.Duration) *barTimeSeriesBuffer {
	return &barTimeSeriesBuffer{
		interval: interval,
		lastAvg:  math.NaN(),
	}
}

const minimumSamplesCount = 5

var _ domain.BarTimeSeriesBuffer = (*barTimeSeriesBuffer)(nil)

type barTimeSeriesBuffer struct {
	interval       time.Duration
	storage        []float64
	lastTruncation int64
	lastAvg        float64
}

func (b *barTimeSeriesBuffer) Add(timestamp int64, value float64) {
	truncation := time.Unix(timestamp, 0).Truncate(b.interval).Unix()

	if truncation != b.lastTruncation {
		// closing bar period
		b.storage = nil
		b.lastTruncation = truncation
		b.lastAvg = b.getAgv() // keep avg from previous interval
	}

	b.storage = append(b.storage, value)
}

func (b *barTimeSeriesBuffer) Get() float64 {
	if len(b.storage) < minimumSamplesCount {
		return b.lastAvg
	}

	return b.getAgv()
}

func (b *barTimeSeriesBuffer) getAgv() float64 {
	avg := 0.0
	for _, v := range b.storage {
		avg += v
	}
	return avg / float64(len(b.storage))
}
