package clock

import "time"

var currentTimezone string = "Asia/Singapore"
var timezone *time.Location = nil

const (
	standardFormat  string = "2006-01-02 15:04:05"
	timestampFormat string = "060102150405"
)

func SetTimezone(newTimezone string) {
	tz, err := time.LoadLocation(newTimezone)
	if err == nil {
		currentTimezone = newTimezone
		timezone = tz
	}
}

func GetTimezone() *time.Location {
	if timezone == nil {
		timezone, _ = time.LoadLocation(currentTimezone)
	}
	return timezone
}

func TimeNow() time.Time {
	return time.Now().In(GetTimezone())
}

func ParseTime(datetime string) (time.Time, error) {
	return time.ParseInLocation(standardFormat, datetime, GetTimezone())
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
