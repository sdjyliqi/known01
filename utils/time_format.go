package utils

import "time"

const (
	DayTime  = "2006-01-02"
	FullTime = "2006-01-02 15:04:05"
)

func ConvertToTime(ts int64) time.Time {
	return time.Unix(ts, 0)
}
