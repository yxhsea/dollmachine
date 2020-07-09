package ff_convert

import (
	"bytes"
	"io"
)

func IoReaderToByte(reader io.Reader) *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf
}
