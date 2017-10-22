package service

import (
	"context"
	"time"

	kitLog "github.com/go-kit/kit/log"
)

// LoggedService = log + service
type LoggedService struct {
	Logger kitLog.Logger
	Next   StringService
}

// Uppercase impl
func (mw LoggedService) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Uppercase(ctx, s)
	return
}

// Count impl
func (mw LoggedService) Count(ctx context.Context, s string) (n int) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.Next.Count(ctx, s)
	return
}
