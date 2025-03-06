package clock

import "time"

var currentTimezone string = "Asia/Singapore"

const (
	standardFormat  string = "2006-01-02 15:04:05"
	timestampFormat string = "060102150405"
)

func SetTimezone(timezone string) {
	_, err := time.LoadLocation(currentTimezone)
	if err == nil {
		currentTimezone = timezone
	}
}

func TimeNow() time.Time {
	timezone, _ := time.LoadLocation(currentTimezone)
	return time.Now().In(timezone)
}

func ParseTime(datetime string) (time.Time, error) {
	timezone, _ := time.LoadLocation(currentTimezone)
	return time.ParseInLocation(standardFormat, datetime, timezone)
}

func DateTimeNow() string {
	return StandardFormat(TimeNow())
}

func TimestampNow() string {
	return TimestampFormat(TimeNow())
}

func StandardFormat(t time.Time) string {
	return t.Format(standardFormat)
}

func TimestampFormat(t time.Time) string {
	return t.Format(timestampFormat)
}

func DateTimeNowWithExpiry(duration time.Duration) (string, string) {
	now := TimeNow()
	expiry := now.Add(duration)
	return StandardFormat(now), StandardFormat(expiry)
}

func CheckIfExpired(expiry string) bool {
	limit, err := ParseTime(expiry)
	if err != nil {
		return true // default to expired if invalid datetime
	}
	return TimeNow().After(limit)
}

func Sleep(pause time.Duration, start time.Time) {
	sleep := pause - TimeNow().Sub(start)
	time.Sleep(sleep)
}
