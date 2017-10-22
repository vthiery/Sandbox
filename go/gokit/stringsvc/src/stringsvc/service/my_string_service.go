package service

import (
	"context"
	"strings"
)

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
