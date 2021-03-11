package utils

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

// NewBackoffUtil :nodoc
func NewBackoffUtil() *backoff.ExponentialBackOff {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 1 * time.Minute
	b.Multiplier = 2
	b.MaxInterval = 15 * time.Minute

	return b
}
