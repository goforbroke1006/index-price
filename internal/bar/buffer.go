package bar

import (
	"math"

	"index-price/domain"
)

func NewRoundBarTimeSeriesBuffer(barType domain.BarType) *barTimeSeriesBuffer {
	return &barTimeSeriesBuffer{barType: barType}
}

var _ domain.BarTimeSeriesBuffer = (*barTimeSeriesBuffer)(nil)

type barTimeSeriesBuffer struct {
	barType        domain.BarType
	storage        []float64
	lastTruncation int64
	lastSnapshot   domain.BarItem
}

func (b *barTimeSeriesBuffer) Add(timestamp int64, value float64) {
	truncation := domain.TruncateWithBarType(b.barType, timestamp)

	if truncation != b.lastTruncation {
		// closing bar period
		b.storage = nil
		b.lastTruncation = truncation
	}

	b.storage = append(b.storage, value)

	b.lastSnapshot.Min = b.getMinPrice()
	b.lastSnapshot.Max = b.getMaxPrice()
	b.lastSnapshot.Open = b.getOpenPrice()
	b.lastSnapshot.Close = b.getClosePrice()
}

func (b *barTimeSeriesBuffer) Get() domain.BarItem {
	return b.lastSnapshot
}

func (b *barTimeSeriesBuffer) getOpenPrice() float64 {
	return b.storage[0]
}

func (b *barTimeSeriesBuffer) getClosePrice() float64 {
	return b.storage[len(b.storage)-1]
}

func (b *barTimeSeriesBuffer) getMinPrice() float64 {
	min := math.MaxFloat64
	for _, v := range b.storage {
		if v < min {
			min = v
		}
	}
	return min
}

func (b *barTimeSeriesBuffer) getMaxPrice() float64 {
	max := 0.0
	for _, v := range b.storage {
		if v > max {
			max = v
		}
	}
	return max
}
