package utils

import "time"

func Uptime(startTime time.Time) time.Duration {
	return time.Since(startTime)
}
