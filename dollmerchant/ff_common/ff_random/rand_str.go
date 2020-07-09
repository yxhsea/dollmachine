package ff_random

import (
	"math/rand"
	"time"
)

//size 长度 ，kind  0:纯数字 1:小写字母  2:大写字母 3:数字、大小写字母
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func KrandNum(size int) string {
	return Krand(size, 0)
}

func KrandLowerChar(size int) string {
	return Krand(size, 1)
}

func KrandUpperChar(size int) string {
	return Krand(size, 2)
}

func KrandAll(size int) string {
	return Krand(size, 3)
}
