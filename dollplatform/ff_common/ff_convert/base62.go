package ff_convert

import (
	"bytes"
	"math"
)

const base62alphabet = "012cduvefg34EFxyBCzAG5IJ67PQRS89abhijklmUVntWXwDopqrsHKLMNOTYZ"

func Base62Encode(number int64) string {
	if number <= 0 {
		return "0"
	}
	chars := make([]byte, 0)
	var length int64 = int64(len(base62alphabet))
	for number > 0 {
		result := number / length
		remainder := number % length
		chars = append(chars, base62alphabet[remainder])
		number = result
	}
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func Base62Decode(token string) int64 {
	var number int64 = 0
	idx := 0.0
	chars := []byte(base62alphabet)
	charsLength := float64(len(chars))
	tokenLength := float64(len(token))
	for _, c := range []byte(token) {
		power := tokenLength - (idx + 1)
		index := int64(bytes.IndexByte(chars, c))
		number += index * int64(math.Pow(charsLength, power))
		idx++
	}
	return number
}
