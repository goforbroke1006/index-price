package bar

import (
	"github.com/stretchr/testify/assert"
	"index-price/domain"
	"testing"
)

func TestRoundBarTimeSeriesBuffer_Add(t *testing.T) {
	t.Parallel()

	t.Run("positive - add measurements inside minute, 2022-08-07 11:05:*", func(t *testing.T) {
		t.Parallel()

		buffer := barTimeSeriesBuffer{barType: domain.BarType1Minute}
		buffer.Add(1659870315, 123.45)
		buffer.Add(1659870349, 122.45)
		buffer.Add(1659870359, 122.45)

		assert.Len(t, buffer.storage, 3)
	})

	t.Run("positive - add measurements outside minute 2022-08-07 11:05:* and 2022-08-07 11:06:*", func(t *testing.T) {
		t.Parallel()

		buffer := barTimeSeriesBuffer{barType: domain.BarType1Minute}
		buffer.Add(1659870315, 123.45) // 05
		buffer.Add(1659870349, 122.45) // 05
		buffer.Add(1659870360, 122.45) // 06
		buffer.Add(1659870365, 122.45) // 06

		assert.Len(t, buffer.storage, 2)
	})

	t.Run("positive - add measurements outside minute 2022-08-07 11:05:* and 2022-08-07 11:06:*", func(t *testing.T) {
		t.Parallel()

		buffer := barTimeSeriesBuffer{barType: domain.BarType5Minute}
		buffer.Add(1659870315, 123.45) // 05
		buffer.Add(1659870349, 122.45) // 05
		buffer.Add(1659870360, 122.45) // 06
		buffer.Add(1659870365, 122.45) // 06

		assert.Len(t, buffer.storage, 4)
	})
}

func TestRoundBarTimeSeriesBuffer_Avg(t *testing.T) {
	t.Parallel()

	t.Run("negative - empty storage", func(t *testing.T) {
		t.Parallel()

		buffer := barTimeSeriesBuffer{barType: domain.BarType1Minute}
		actual := buffer.Avg()
		expect := 0.0
		assert.Equal(t, expect, actual)
	})

	t.Run("positive - add measurements outside minute 2022-08-07 11:05:* and 2022-08-07 11:06:*", func(t *testing.T) {
		t.Parallel()

		buffer := barTimeSeriesBuffer{barType: domain.BarType1Minute}
		buffer.Add(1659870315, 123.45) // 05
		buffer.Add(1659870349, 122.45) // 05
		buffer.Add(1659870360, 121.45) // 06
		buffer.Add(1659870365, 120.45) // 06

		actual := buffer.Avg()
		expect := (121.45 + 120.45) / 2
		assert.Equal(t, expect, actual)
	})
}
