package bar

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"index-price/domain"
)

func TestRoundBarTimeSeriesBuffer_Add_Get(t *testing.T) {
	t.Parallel()

	t.Run("positive - add measurements inside minute, 2022-08-07 11:05:*", func(t *testing.T) {
		t.Parallel()

		buffer := barTimeSeriesBuffer{barType: domain.BarType1Minute}

		{
			buffer.Add(1659870315, 123.45)
			item := buffer.Get()

			//assert.True(t, item.Empty())
			assert.Equal(t, 123.45, item.Open)
			assert.Equal(t, 123.45, item.Close)
			assert.Equal(t, 123.45, item.Max)
			assert.Equal(t, 123.45, item.Min)
		}

		{
			buffer.Add(1659870349, 122.45)
			item := buffer.Get()

			//assert.False(t, item.Empty())
			assert.Equal(t, 123.45, item.Open)
			assert.Equal(t, 122.45, item.Close)
			assert.Equal(t, 123.45, item.Max)
			assert.Equal(t, 122.45, item.Min)
		}

		{
			buffer.Add(1659870359, 121.45)
			item := buffer.Get()

			//assert.False(t, item.Empty())
			assert.Equal(t, 123.45, item.Open)
			assert.Equal(t, 121.45, item.Close)
			assert.Equal(t, 123.45, item.Max)
			assert.Equal(t, 121.45, item.Min)
		}

		assert.Len(t, buffer.storage, 3)
	})

	t.Run("positive - add measurements outside minute 2022-08-07 11:05:* and 2022-08-07 11:06:*", func(t *testing.T) {
		t.Parallel()

		buffer := barTimeSeriesBuffer{barType: domain.BarType1Minute}
		{
			buffer.Add(1659870315, 123.45) // 05

			item := buffer.Get()
			assert.Equal(t, 123.45, item.Open)
			assert.Equal(t, 123.45, item.Close)
			assert.Equal(t, 123.45, item.Max)
			assert.Equal(t, 123.45, item.Min)
		}

		{
			buffer.Add(1659870349, 122.45) // 05

			item := buffer.Get()
			assert.Equal(t, 123.45, item.Open)
			assert.Equal(t, 122.45, item.Close)
			assert.Equal(t, 123.45, item.Max)
			assert.Equal(t, 122.45, item.Min)
		}

		{
			buffer.Add(1659870360, 121.45) // 06

			item := buffer.Get()
			assert.Equal(t, 121.45, item.Open)
			assert.Equal(t, 121.45, item.Close)
			assert.Equal(t, 121.45, item.Max)
			assert.Equal(t, 121.45, item.Min)
		}

		{
			buffer.Add(1659870365, 120.45) // 06

			item := buffer.Get()
			assert.Equal(t, 121.45, item.Open)
			assert.Equal(t, 120.45, item.Close)
			assert.Equal(t, 121.45, item.Max)
			assert.Equal(t, 120.45, item.Min)
		}

		assert.Len(t, buffer.storage, 2)
	})

	t.Run("positive - add measurements outside minute 2022-08-07 11:05:* and 2022-08-07 11:06:*", func(t *testing.T) {
		t.Parallel()

		buffer := barTimeSeriesBuffer{barType: domain.BarType5Minute}
		buffer.Add(1659870315, 123.45) // 05
		buffer.Add(1659870349, 127.45) // 05
		buffer.Add(1659870360, 120.45) // 06
		buffer.Add(1659870365, 125.45) // 06

		item := buffer.Get()
		assert.Equal(t, 123.45, item.Open)
		assert.Equal(t, 125.45, item.Close)
		assert.Equal(t, 127.45, item.Max)
		assert.Equal(t, 120.45, item.Min)

		assert.Len(t, buffer.storage, 4)
	})
}
