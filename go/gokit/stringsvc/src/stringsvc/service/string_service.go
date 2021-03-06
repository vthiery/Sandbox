package service

import (
	"context"
)

// StringService interface
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}
