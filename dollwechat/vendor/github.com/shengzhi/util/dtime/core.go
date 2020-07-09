// Package dtime 工具包 时间辅助操作
package dtime

import (
	"fmt"
	"time"
)

// JSONShortTime 只展示年月日 e.g. 2017--11-01
type JSONShortTime int64

// MarshalJSON outputs JSON presentation
func (t JSONShortTime) MarshalJSON() ([]byte, error) {
	if int64(t) <= 0 {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, time.Unix(int64(t), 0).Format("2006-01-02"))), nil
}

// UnmarshalJSON unmarshal string to JSONTime
func (t *JSONShortTime) UnmarshalJSON(b []byte) error {
	if len(b) <= 0 || string(b) == `""` {
		*t = 0
		return nil
	}
	tm, err := time.ParseInLocation(`"2006-01-02"`, string(b), time.Local)
	if err != nil {
		return err
	}
	*t = JSONShortTime(tm.Unix())
	return nil
}

// JSONMiddleTime  展示年月日时分,e.g. 2017-11-01 14:23
type JSONMiddleTime int64

func (t JSONMiddleTime) String() string {
	if t == 0 {
		return ""
	}
	return time.Unix(int64(t), 0).Format("2006-01-02 15:04")
}

// MarshalJSON outputs JSON presentation
func (t JSONMiddleTime) MarshalJSON() ([]byte, error) {
	if int64(t) <= 0 {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, time.Unix(int64(t), 0).Format("2006-01-02 15:04"))), nil
}

// UnmarshalJSON unmarshal string to JSONTime
func (t *JSONMiddleTime) UnmarshalJSON(b []byte) error {
	if len(b) <= 0 || string(b) == `""` {
		*t = 0
		return nil
	}
	tm, err := time.ParseInLocation(`"2006-01-02 15:04"`, string(b), time.Local)
	if err != nil {
		return err
	}
	*t = JSONMiddleTime(tm.Unix())
	return nil
}

// JSONTime JSON 时间， 时间戳
type JSONTime int64

// MarshalJSON outputs JSON presentation
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if int64(t) <= 0 {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, time.Unix(int64(t), 0).Format("2006-01-02 15:04:05"))), nil
}

// UnmarshalJSON unmarshal string to JSONTime
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	if len(b) <= 0 || string(b) == `""` {
		*t = 0
		return nil
	}
	tm, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), time.Local)
	if err != nil {
		return err
	}
	*t = JSONTime(tm.Unix())
	return nil
}

func (t JSONTime) Time() time.Time {
	return time.Unix(int64(t), 0)
}
func Now() JSONTime { return JSONTime(time.Now().Unix()) }

func Today() JSONTime {
	y, m, d := time.Now().Date()
	return JSONTime(time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix())
}
