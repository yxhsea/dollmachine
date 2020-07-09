package cut_time

import (
	"fmt"
	"time"
)

func GetDayTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetYear() int {
	t := time.Now()
	return t.Year()
}

func GetYearMonth() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d",
		t.Year(), t.Month())
}

func GetDayTime2() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())
}

func GetDay(format string) string {
	return time.Now().Format(format)
}

func GetTimeStamp() int64 {
	return time.Now().Unix()
}

// convert time to timestamp
// t params like "2006-01-02 15:04:05"
// return int64
func StrToTS(t string) int64 {
	loc, _ := time.LoadLocation("Local")
	s, err := time.ParseInLocation("2006-01-02 15:04:05", t, loc)
	if err != nil {
		return 0
	}
	return s.Unix()
}