package clock

import (
	"fmt"
	"strings"
	"time"
	_ "time/tzdata" // include the timezone data package
)

// Default timezone: Asia/Singapore
var currentTimezone string = "Asia/Singapore"
var timezone *time.Location = nil

// Override the current timezone
func SetTimezone(newTimezone string) error {
	tz, err := time.LoadLocation(newTimezone)
	if err == nil {
		currentTimezone = newTimezone
		timezone = tz
	}
	return err
}

// Get current timezone
func GetTimezone() *time.Location {
	if timezone == nil {
		timezone, _ = time.LoadLocation(currentTimezone)
	}
	return timezone
}

// Returns the current time in current timezone
func TimeNow() time.Time {
	return time.Now().In(GetTimezone())
}

// Check if hour/minute now is 00:00 (midnight)
func IsMidnight() bool {
	return HourMinNow() == "00:00"
}

// Extend the given datetime with the given duration
// Return the extended time in standard format (yyyy-mm-dd hh:mm:ss)
func ExtendTime(datetime string, duration time.Duration) string {
	t, err := ParseTime(datetime)
	if err != nil {
		// If given datetime is invalid, return it as it is
		return datetime
	}
	return StandardFormat(t.Add(duration))
}

// Check if given datetime is already expired (before current time)
func CheckIfExpired(expiry string) bool {
	limit, err := ParseTime(expiry)
	if err != nil {
		// Default to expired if invalid datetime
		return true
	}
	return TimeNow().After(limit)
}

// Calculate the duration since the given datetime, rounded to the given duration
// Returns the string format of the duration
func DurationSince(datetime string, round time.Duration) (string, error) {
	t, err := ParseTime(datetime)
	if err != nil {
		return "", err
	}
	duration := TimeNow().Sub(t)
	duration = duration.Round(round)
	return fmt.Sprintf("%v", duration), nil
}

// Check if given yyyy-mm-dd hh:mm:ss is valid
func IsValidDateTime(datetime string) bool {
	datetime = strings.TrimSpace(datetime)
	if datetime == "" {
		return false
	}
	_, err := time.Parse(standardFormat, datetime)
	return err == nil
}
