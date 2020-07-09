package ff_convert

import "strconv"

func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func IntToStr(num int) string {
	return strconv.Itoa(num)
}
