package clock

import (
	"fmt"
	"strings"
	"time"
)

// Return current date in yyyy-mm-dd format
func DateNow() string {
	return DateFormat(Now())
}

// Attach 00:00:00 to given date
func DateTimeStart(date string) string {
	return fmt.Sprintf("%s 00:00:00", date)
}

// Attach 23:59:59 to given date
func DateTimeEnd(date string) string {
	return fmt.Sprintf("%s 23:59:59", date)
}

// Return midnight datetime today in standard format (yyyy-mm-dd hh:mm:ss)
func MidnightToday() string {
	return DateTimeStart(DateNow())
}

// Check if given yyyy-mm-dd date is valid
func IsValidDate(date string) bool {
	_, err := time.Parse(dateFmt, strings.TrimSpace(date))
	return err == nil
}

// Extracts date from datetime string
func ExtractDate(datetime string) string {
	if !IsValidDateTime(datetime) {
		return ""
	}
	return strings.Fields(datetime)[0]
}
