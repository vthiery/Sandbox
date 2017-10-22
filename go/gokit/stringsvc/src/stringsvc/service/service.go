package service

import (
	"context"
	"errors"
	"strings"
)

// StringService interface
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

// MyStringService implements the StringService interface
type MyStringService struct{}

// Uppercase impl
func (MyStringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// Count impl
func (MyStringService) Count(_ context.Context, s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("String is empty")
