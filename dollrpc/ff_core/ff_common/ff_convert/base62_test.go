package ff_convert

import (
	"testing"
)

func TestBase62Encode(t *testing.T) {
	t.Logf("Base62Encode(0) == 0, encode:%s, rs:%t", Base62Encode(0), Base62Encode(0) == "0")
	var tint int64
	var tstr string

	tint = 99
	tstr = "1k"
	t.Logf("int:%d, encode:%s, rs:%t", tint, Base62Encode(tint), Base62Encode(tint) == tstr)
	t.Logf("str:%s, decode:%d, rs:%t", tstr, Base62Decode(tstr), Base62Decode(tstr) == tint)

	tint = 9999999
	tstr = "VTRA"
	t.Logf("int:%d, encode:%s, rs:%t", tint, Base62Encode(tint), Base62Encode(tint) == tstr)
	t.Logf("str:%s, decode:%d, rs:%t", tstr, Base62Decode(tstr), Base62Decode(tstr) == tint)

	tint = 112412401012002
	tstr = "9Nu3Pt9d"
	t.Logf("int:%d, encode:%s, rs:%t", tint, Base62Encode(tint), Base62Encode(tint) == tstr)
	t.Logf("str:%s, decode:%d, rs:%t", tstr, Base62Decode(tstr), Base62Decode(tstr) == tint)
}
