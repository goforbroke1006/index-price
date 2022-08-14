package bar

import (
	"sort"
	"sync"
	"time"

	"index-price/domain"
)

func NewSamplesBuffer(interval time.Duration) *samplesBuffer {
	b := &samplesBuffer{
		interval: interval,
		storage:  make(map[int64][]float64),
	}

	go func() {
		for {
			<-time.After(time.Minute)
			b.clear()
		}
	}()

	return b
}

var _ domain.TickPriceSamplesBuffer = (*samplesBuffer)(nil)

type samplesBuffer struct {
	interval time.Duration

	storage   map[int64][]float64
	storageMx sync.RWMutex
}

func (b *samplesBuffer) Add(timestamp int64, value float64) {
	b.storageMx.Lock()
	defer b.storageMx.Unlock()

	truncation := b.getKey(timestamp)
	b.storage[truncation] = append(b.storage[truncation], value)
}

func (b *samplesBuffer) Get(timestamp int64) float64 {
	b.storageMx.RLock()
	defer b.storageMx.RUnlock()

	truncation := b.getKey(timestamp)
	samples := b.storage[truncation]

	avg := 0.0
	for _, v := range samples {
		avg += v
	}
	return avg / float64(len(samples))
}

func (b *samplesBuffer) getKey(timestamp int64) int64 {
	return time.Unix(timestamp, 0).Truncate(b.interval).Unix()
}

func (b *samplesBuffer) clear() {
	b.storageMx.Lock()
	defer b.storageMx.Unlock()

	const maxStorageLen = 3
	if len(b.storage) <= maxStorageLen {
		return
	}

	extra := len(b.storage) - maxStorageLen

	keys := make([]int64, 0, len(b.storage))
	for k := range b.storage {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for i := 0; i < extra; i++ {
		delete(b.storage, keys[i])
	}
}
