package ff_hash

import (
	"crypto/md5"
	"encoding/hex"
	"dollmachine/dollplatform/ff_common/ff_convert"
)

func MD5Encode(data string) []byte {
	h := md5.New()
	h.Write(ff_convert.Str2Byte(data))
	return h.Sum(nil)
}

func MD5EncodeToString(data string) string {
	y := make([]byte, 32)
	hex.Encode(y, MD5Encode(data))
	return ff_convert.Byte2Str(y)
}
