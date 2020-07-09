package mercury

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/codec"
)

type mercuryCodec struct {
	sync.Mutex
	rwc io.ReadWriteCloser
	mt  codec.MessageType
	buf *bytes.Buffer
}

func (c *mercuryCodec) Close() error {
	c.buf.Reset()
	return c.rwc.Close()
}

func (c *mercuryCodec) String() string {
	return "mercury"
}

func (c *mercuryCodec) Write(m *codec.Message, b interface{}) error {
	switch m.Type {
	case codec.Request:
		data, err := proto.Marshal(b.(proto.Message))
		if err != nil {
			return err
		}
		c.rwc.Write(data)
		m.Header["Content-Encoding"] = "request"
		m.Header["Service"] = m.Target
		m.Header["Endpoint"] = m.Method
	case codec.Response:
		m.Header["Content-Encoding"] = "response"
		data, err := proto.Marshal(b.(proto.Message))
		if err != nil {
			return err
		}
		c.rwc.Write(data)
	case codec.Publication:
		data, err := proto.Marshal(b.(proto.Message))
		if err != nil {
			return err
		}
		c.rwc.Write(data)
	default:
		return fmt.Errorf("Unrecognised message type: %v", m.Type)
	}
	return nil
}

func (c *mercuryCodec) ReadHeader(m *codec.Message, mt codec.MessageType) error {
	c.buf.Reset()
	c.mt = mt

	switch mt {
	case codec.Request:
		m.Method = m.Header["Endpoint"]
		io.Copy(c.buf, c.rwc)
	case codec.Response:
		io.Copy(c.buf, c.rwc)
	case codec.Publication:
		io.Copy(c.buf, c.rwc)
	default:
		return fmt.Errorf("Unrecognised message type: %v", mt)
	}
	return nil
}

func (c *mercuryCodec) ReadBody(b interface{}) error {
	var data []byte
	switch c.mt {
	case codec.Request, codec.Response:
		data = c.buf.Bytes()
	case codec.Publication:
		data = c.buf.Bytes()
	default:
		return fmt.Errorf("Unrecognised message type: %v", c.mt)
	}
	if b != nil {
		return proto.Unmarshal(data, b.(proto.Message))
	}
	return nil
}

func NewCodec(rwc io.ReadWriteCloser) codec.Codec {
	return &mercuryCodec{
		buf: bytes.NewBuffer(nil),
		rwc: rwc,
	}
}
