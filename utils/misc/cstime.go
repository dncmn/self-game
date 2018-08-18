package misc

import (
	"time"
)

var (
	cstZone *time.Location
	offset  time.Duration
)

func init() {
	cstZone, _ = time.LoadLocation("Asia/Shanghai")
}

func SetLocation(location string) (err error) {
	cstZone, err = time.LoadLocation(location)
	return err
}

func GetTimestamp() int64 {
	return time.Now().In(cstZone).Add(offset).Unix()
}

func GetTime() time.Time {
	return time.Now().In(cstZone).Add(offset)
}

func UnixMilli() int64 {
	t := GetTime()
	return t.Round(time.Millisecond).UnixNano() /
		(int64(time.Millisecond) / int64(time.Nanosecond))
}

func FromUnixMilli(m int64) time.Time {
	return time.Unix(m/1e3, (m%1e3)*int64(time.Millisecond)/int64(time.Nanosecond))
}

func ToSqlTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetSqlTimestamp() string {
	return GetTime().Format("2006-01-02 15:04:05")
}

func FromUnixTime(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec).In(cstZone)
}

func FromDate(year int, month time.Month, day, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, cstZone)
}

func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, cstZone)
}

func SinceTime(t time.Time) time.Duration {
	return GetTime().Sub(t)
}

func SetTimeOffset(s string) error {
	var err error
	if s == "0" {
		offset, _ = time.ParseDuration("0")
		return err
	}
	offsetAdj, err := time.ParseDuration(s)
	if err != nil {
		offset, _ = time.ParseDuration("0")
	}
	offset += offsetAdj
	return err
}

func GetTimeOffset() string {
	return offset.String()
}

func GetTimeZone() *time.Location {
	return cstZone
}
