package clock

import "time"

const (
	timeFormat      string = "15:04:05"
	hourMinsFormat  string = "15:04"
	standardFormat  string = "2006-01-02 15:04:05"
	timestampFormat string = "060102150405"
)

// Return current datetime in standard format (yyyy-mm-dd hh:mm:ss)
func DateTimeNow() string {
	return StandardFormat(TimeNow())
}

// Return datetime now and an expiry datetime (based on given duration),
// Both in standard format (yyyy-mm-dd hh:mm:ss)
func DateTimeNowWithExpiry(duration time.Duration) (now string, expiry string) {
	timeNow := TimeNow()
	timeExpiry := timeNow.Add(duration)
	now, expiry = StandardFormat(timeNow), StandardFormat(timeExpiry)
	return now, expiry
}

// Return current datetime in timestamp format (yymmddhhmmss)
func TimestampNow() string {
	return TimestampFormat(TimeNow())
}

// Return current hour & minute in hh:mm format
func HourMinNow() string {
	return TimeNow().Format(hourMinsFormat)
}

// Format the given time in standard datetime format (yyyy-mm-dd hh:mm:ss)
func StandardFormat(t time.Time) string {
	return t.Format(standardFormat)
}

// Format the given time in timestamp format (yymmddhhmmss)
func TimestampFormat(t time.Time) string {
	return t.Format(timestampFormat)
}

// Format the given time in time format (hh:mm:ss)
func TimeFormat(t time.Time) string {
	return t.Format(timeFormat)
}

// Parse the given string as a datetime in standard format, in current timezone
func ParseTime(datetime string) (time.Time, error) {
	return time.ParseInLocation(standardFormat, datetime, GetTimezone())
}
