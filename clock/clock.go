// Package clock contains time and date-related functions.
package clock

import (
	"time"
	_ "time/tzdata" // include the timezone data package
)

// Default timezone: Asia/Singapore
var currentTimezone string = "Asia/Singapore"
var timezone *time.Location = nil

// Override current timezone
func SetTimezone(newTimezone string) error {
	tz, err := time.LoadLocation(newTimezone)
	if err == nil {
		currentTimezone = newTimezone
		timezone = tz
	}
	return err
}

// Get current timezone location object
func CurrentTimezone() *time.Location {
	if timezone == nil {
		timezone, _ = time.LoadLocation(currentTimezone)
	}
	return timezone
}

// Pauses for given duration, taking into account
// the starting time, adjusting the pause duration
func Sleep(pause time.Duration, start time.Time) {
	// subtract time elapsed since start
	duration := pause - Now().Sub(start)
	time.Sleep(duration)
}

// Run infinite loop; used for main() with goroutines
func RunForever() {
	for {
		select {}
	}
}
