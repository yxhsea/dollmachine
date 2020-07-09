package helper

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// GetExecDir 获取当前程序运行目录
func GetExecDir() (dir, file string) {
	pwd, _ := exec.LookPath(os.Args[0])
	_, file = filepath.Split(pwd)
	dir = filepath.Dir(pwd)
	return
}

const Digits = "0123456789"
const DigitsAndLetters = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"

// RandomGenerator 随机字符串生成器
type RandomGenerator func(int) string

// CreateRandomGenerator 创建随机字符生成器
func CreateRandomGenerator(pool string) RandomGenerator {
	return func(length int) string {
		rand.Seed(time.Now().Unix())
		code := make([]byte, length)
		for i := 0; i < length; i++ {
			code[i] = pool[rand.Intn(len(pool)-1)]
		}
		return string(code)
	}

}

// IPNumber 数字IP
type IPNumber int64

func (ip IPNumber) MarshalJSON() ([]byte, error) {
	return []byte(`"` + Int2IP(int64(ip)) + `"`), nil
}

// IP2Int IP地址转换为数字
func IP2Int(ip string) int64 {
	ipv4 := net.ParseIP(ip)
	ipint := big.NewInt(0)
	ipint.SetBytes(ipv4.To4())
	return ipint.Int64()
}

// Int2IP 数字转换为IP V4
func Int2IP(ipint int64) string {
	var bytes [4]byte
	bytes[0] = byte(ipint & 0xFF)
	bytes[1] = byte((ipint >> 8) & 0xFF)
	bytes[2] = byte((ipint >> 16) & 0xFF)
	bytes[3] = byte((ipint >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String()
}

// Joins 根据指定连接字符串拼接数组
func Joins(slice []interface{}, delim string) string {
	if len(slice) <= 0 {
		return ""
	}
	if len(slice) == 1 {
		return fmt.Sprintf("%v", slice[0])
	}
	var buf bytes.Buffer
	for i := 0; i < len(slice)-1; i++ {
		buf.WriteString(fmt.Sprintf("%v", slice[i]))
		buf.WriteString(delim)
	}
	buf.WriteString(fmt.Sprintf("%v", slice[len(slice)-1]))
	return buf.String()
}
