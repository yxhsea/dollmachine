package ff_v1flow

import (
	"fmt"
	"os"
	"testing"
	//"time"
)

func TestName(t *testing.T) {
	//merchantId := "20000332793"
	/*date := time.Now().Format("2006-01-02")*/
	dir := fmt.Sprintf("./upload/20000332793/2006-01-02")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
