package bar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const floatDelta = 0.001

func TestRoundBarTimeSeriesBuffer_Add_Get(t *testing.T) {
	t.Parallel()

	t.Run("positive - single interval, single measurement", func(t *testing.T) {
		t.Parallel()

		buffer := samplesBuffer{
			interval: 5 * time.Second,
			storage:  map[int64][]float64{},
		}

		{
			buffer.Add(1659870315, 123.40) // Sun Aug 07 2022 11:05:15 GMT+0000
			assert.Equal(t, 123.40, buffer.Get(1659870315))
		}
	})

	t.Run("positive - single interval, multiple measurements", func(t *testing.T) {
		t.Parallel()

		buffer := samplesBuffer{
			interval: 5 * time.Second,
			storage:  map[int64][]float64{},
		}

		{
			buffer.Add(1659870315, 123.40) // Sun Aug 07 2022 11:05:15 GMT+0000
			assert.InDelta(t, 123.40, buffer.Get(1659870315), floatDelta)
		}
		{
			buffer.Add(1659870316, 123.40) // Sun Aug 07 2022 11:05:15 GMT+0000
			assert.InDelta(t, 123.40, buffer.Get(1659870315), floatDelta)
		}
		{
			buffer.Add(1659870316, 123.40) // Sun Aug 07 2022 11:05:15 GMT+0000
			assert.InDelta(t, 123.40, buffer.Get(1659870315), floatDelta)
		}
	})

	t.Run("positive - add measurements inside minute, 2022-08-07 11:05:*", func(t *testing.T) {
		t.Parallel()

		buffer := samplesBuffer{
			interval: 5 * time.Second,
			storage:  map[int64][]float64{},
		}

		{
			buffer.Add(1659870315, 123.40) // Sun Aug 07 2022 11:05:15 GMT+0000
			assert.InDelta(t, 123.40, buffer.Get(1659870315), floatDelta)
		}

		{
			buffer.Add(1659870319, 123.60)
			assert.InDelta(t, 123.50, buffer.Get(1659870315), floatDelta)
		}

		// next time interval

		{
			buffer.Add(1659870321, 120.72) // Sun Aug 07 2022 11:05:15 GMT+0000
			assert.InDelta(t, 120.72, buffer.Get(1659870320), floatDelta)
		}

		{
			buffer.Add(1659870322, 120.42) // Sun Aug 07 2022 11:05:15 GMT+0000
			assert.InDelta(t, 120.57, buffer.Get(1659870320), floatDelta)
		}

		{
			buffer.Add(1659870324, 110.42) // Sun Aug 07 2022 11:05:15 GMT+0000
			assert.InDelta(t, 117.18666666666667, buffer.Get(1659870320), floatDelta)
		}
	})
}
