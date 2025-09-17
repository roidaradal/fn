package clock

import (
	"fmt"
	"strings"
	"time"
	_ "time/tzdata"
)

var currentTimezone string = "Asia/Singapore"
var timezone *time.Location = nil

const (
	dateFormat      string = "2006-01-02"
	timeFormat      string = "15:04:05"
	hourMinsFormat  string = "15:04"
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

func DateNow() string {
	return DateFormat(TimeNow())
}

func MidnightToday() string {
	return DateNow() + " 00:00:00"
}

func IsMidnight() bool {
	return HourMinNow() == "00:00"
}

func HourMinNow() string {
	return TimeNow().Format(hourMinsFormat)
}

func TimestampNow() string {
	return TimestampFormat(TimeNow())
}

func DateFormat(t time.Time) string {
	return t.Format(dateFormat)
}

func StandardFormat(t time.Time) string {
	return t.Format(standardFormat)
}

func TimestampFormat(t time.Time) string {
	return t.Format(timestampFormat)
}

func TimeFormat(t time.Time) string {
	return t.Format(timeFormat)
}

func DateTimeNowWithExpiry(duration time.Duration) (string, string) {
	now := TimeNow()
	expiry := now.Add(duration)
	return StandardFormat(now), StandardFormat(expiry)
}

func ExtendTime(datetime string, duration time.Duration) string {
	t, err := ParseTime(datetime)
	if err != nil {
		return datetime
	}
	return StandardFormat(t.Add(duration))
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

func DurationSince(datetime string, round time.Duration) (string, error) {
	t, err := ParseTime(datetime)
	if err != nil {
		return "", err
	}
	duration := TimeNow().Sub(t)
	duration = duration.Round(round)
	return fmt.Sprintf("%v", duration), nil
}

func IsValidDate(date string) bool {
	date = strings.TrimSpace(date)
	if date == "" {
		return false
	}
	_, err := time.Parse(dateFormat, date)
	return err == nil
}

func DateStart(date string) string {
	return fmt.Sprintf("%s 00:00:00", date)
}

func DateEnd(date string) string {
	return fmt.Sprintf("%s 23:59:59", date)
}
