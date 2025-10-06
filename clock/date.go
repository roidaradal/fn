package clock

import (
	"fmt"
	"strings"
	"time"
)

const dateFormat string = "2006-01-02"

// Return current date in yyyy-mm-dd format
func DateNow() string {
	return DateFormat(TimeNow())
}

// Format the given time in date format (yyyy-mm-dd)
func DateFormat(t time.Time) string {
	return t.Format(dateFormat)
}

// Check if given yyyy-mm-dd date is valid
func IsValidDate(date string) bool {
	date = strings.TrimSpace(date)
	if date == "" {
		return false
	}
	_, err := time.Parse(dateFormat, date)
	return err == nil
}

// Attach 00:00:00 to given date (first second)
func DateStart(date string) string {
	return fmt.Sprintf("%s 00:00:00", date)
}

// Attach 23:59:59 to given date (last second)
func DateEnd(date string) string {
	return fmt.Sprintf("%s 23:59:59", date)
}

// Return the midnight datetime today in standard format (yyyy-mm-dd hh:mm:ss)
func MidnightToday() string {
	return fmt.Sprintf("%s 00:00:00", DateNow())
}
