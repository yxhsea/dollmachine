package ff_time

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {

	mstZone, err := time.LoadLocation("Asia/Chongqing")
	fmt.Println(mstZone, err, err)
	n := time.Now()
	if fmt.Sprintf("%v", n) == fmt.Sprintf("%v", n.In(mstZone)) {
		t.Logf("you are testing and in CST %v", mstZone)
	}

	zeroTime := time.Time{}.Unix()
	ts, err := ParseAny("INVALID", mstZone)
	fmt.Println(ts, err, zeroTime)

	ts, err = ParseAny("2017-04-01 22:43:22 +0800 UTC", mstZone)
	fmt.Println(ts.In(mstZone).Unix(), err)
	fmt.Println(ts.Zone())

	ts, err = ParseAny("2017-04-01 22:43:22", mstZone)
	fmt.Println(ts.Unix(), err)

	ts, err = ParseAny("2014-04-26 17:24:37.123", mstZone)
	fmt.Println(ts.Unix(), err)

	ts, err = ParseAny("2014-04-26", mstZone)
	fmt.Println(ts.Unix(), err)

	ts, err = ParseAny("1332151919", mstZone)
	fmt.Println(ts.Unix(), err)

	ts, err = ParseAny("1384216367189", mstZone)
	fmt.Println(ts.Unix(), err)

}
