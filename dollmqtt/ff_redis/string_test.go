package ff_redis

import (
	"testing"
	"fmt"
	"log"
	"runtime"
)

func TestNewString(t *testing.T) {
	str := "'%v' abc"
	fmt.Println(fmt.Sprintf(str,""))
}

func TestString_Del(t *testing.T) {

}

func TestNewString2(t *testing.T) {
	log.Println(filePath(), "Error : this is err")
}

func filePath() string {
	_,file,line,_ := runtime.Caller(1)
	return file + ":" + fmt.Sprint(line)
}