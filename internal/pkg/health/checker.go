package health

import (
	"time"

	"github.com/alexliesenfeld/health"
)

func NewChecker(checkers ...health.CheckerOption) health.Checker {
	checks := []health.CheckerOption{
		// Set the time-to-live for our cache to 1 second (default).
		health.WithCacheDuration(1 * time.Second),

		// Configure a global timeout that will be applied to all checks.
		health.WithTimeout(10 * time.Second),
	}

	checks = append(checks, checkers...)

	return health.NewChecker(
		checks...,
	)
}
