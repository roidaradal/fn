package clock

import (
	"fmt"
	"strings"
	"time"
)

// Returns current time in current timezone
func Now() time.Time {
	return time.Now().In(CurrentTimezone())
}

// Return current datetime in standard format (yyyy-mm-dd hh:mm:ss)
func DateTimeNow() string {
	return StandardFormat(Now())
}

// Return datetime now and expiry datetime (based on given duration),
// Both in standard format (yyyy-mm-dd hh:mm:ss)
func DateTimeNowWithExpiry(duration time.Duration) (now, expiry string) {
	timeNow := Now()
	timeExpiry := timeNow.Add(duration)
	now, expiry = StandardFormat(timeNow), StandardFormat(timeExpiry)
	return now, expiry
}

// Return current time in hh:mm:ss format
func TimeNow() string {
	return TimeFormat(Now())
}

// Return current hour and minute in hh:mm format
func HourMinNow() string {
	return HourMinFormat(Now())
}

// Return current datetime in timestamp format (yymmddhhmmss)
func TimestampNow() string {
	return TimestampFormat(Now())
}

// Parse given string as datetime in standard format, in current timezone
func ParseDateTime(datetime string) (time.Time, error) {
	datetime = strings.TrimSpace(datetime)
	return time.ParseInLocation(standardFmt, datetime, CurrentTimezone())
}

// Check if given datetime in standard format is valid
func IsValidDateTime(datetime string) bool {
	_, err := time.Parse(standardFmt, strings.TrimSpace(datetime))
	return err == nil
}

// Check if hour:minute now is 00:00 (midnight)
func IsMidnight() bool {
	return HourMinNow() == "00:00"
}

// Check if given datetime is already expired (before current time)
func IsExpired(expiry string) bool {
	limit, err := ParseDateTime(expiry)
	if err != nil {
		// Default to expired if invalid datetime
		return true
	}
	return Now().After(limit)
}

// Extend given datetime with given duration,
// Return extended time in standard format (yyyy-mm-dd hh:mm:ss)
func ExtendTime(datetime string, duration time.Duration) (string, error) {
	dt, err := ParseDateTime(datetime)
	if err != nil {
		return "", err
	}
	return StandardFormat(dt.Add(duration)), nil
}

// Calculate duration since given datetime, rounded to given duration
// Return duration as string
func DurationSince(datetime string, round time.Duration) (string, error) {
	dt, err := ParseDateTime(datetime)
	if err != nil {
		return "", err
	}
	duration := Now().Sub(dt).Round(round)
	return fmt.Sprintf("%v", duration), nil
}
