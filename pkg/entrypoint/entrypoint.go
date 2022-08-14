package entrypoint

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func NewSignalContext(ctx context.Context) context.Context {
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	return ctx
}

func WaitForShutdown(ctx context.Context) {
	<-ctx.Done()
	time.Sleep(time.Second)
}
