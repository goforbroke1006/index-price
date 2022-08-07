package bar

import "index-price/domain"

func NewRoundBarTimeSeriesBuffer(barType domain.BarType) *barTimeSeriesBuffer {
	return &barTimeSeriesBuffer{
		barType: barType,
	}
}

var _ domain.BarTimeSeriesBuffer = (*barTimeSeriesBuffer)(nil)

type barTimeSeriesBuffer struct {
	barType        domain.BarType
	storage        []float64
	lastTruncation int64
}

func (b *barTimeSeriesBuffer) Add(timestamp int64, value float64) {
	truncation := domain.TruncateWithBarType(b.barType, timestamp)

	if truncation != b.lastTruncation {
		b.storage = nil
		b.lastTruncation = truncation
	}

	b.storage = append(b.storage, value)
}

func (b barTimeSeriesBuffer) Avg() float64 {
	if len(b.storage) == 0 {
		return 0
	}

	var sum float64

	for _, val := range b.storage {
		sum += val
	}

	return sum / float64(len(b.storage))
}
