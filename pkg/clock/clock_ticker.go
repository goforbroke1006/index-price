package clock

import (
	"time"
)

func NewOnClockTicker(step time.Duration) time.Ticker {
	c := make(chan time.Time, 1)

	go func() {
		for {
			next := time.Now().Truncate(step).Add(step)
			//zap.L().Info("next", zap.String("tick", next.Format(time.RFC3339)))
			waiting := time.Until(next) // next.Sub(time.Now())
			<-time.After(waiting)

			c <- time.Now()
		}
	}()

	t := time.Ticker{C: c}
	return t
}
