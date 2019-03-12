package time

import (
	"time"
)

var startTime = time.Unix(0, 0)

// DayID returns day number from unix
func DayID(from time.Time) int {
	return int(from.Sub(startTime).Hours() / 24)
}

// DaysBetween do smth
func DaysBetween(from, to time.Time) int {
	return DayID(to) - DayID(from)
}

// StartOfDay returns start of the day
func StartOfDay(from time.Time, location *time.Location) time.Time {
	if location == nil {
		location = from.Location()
	}
	return time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, location)
}
