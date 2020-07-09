package rabbitmq

import (
	"crypto/tls"
	"errors"
	"testing"

	"github.com/streadway/amqp"
)

func TestNewRabbitMQConnURL(t *testing.T) {
	testcases := []struct {
		title string
		urls  []string
		want  string
	}{
		{"Multiple URLs", []string{"amqp://example.com/one", "amqp://example.com/two"}, "amqp://example.com/one"},
		{"Insecure URL", []string{"amqp://example.com"}, "amqp://example.com"},
		{"Secure URL", []string{"amqps://example.com"}, "amqps://example.com"},
		{"Invalid URL", []string{"http://example.com"}, DefaultRabbitURL},
		{"No URLs", []string{}, DefaultRabbitURL},
	}

	for _, test := range testcases {
		conn := newRabbitMQConn("exchange", test.urls)

		if have, want := conn.url, test.want; have != want {
			t.Errorf("%s: invalid url, want %q, have %q", test.title, want, have)
		}
	}
}

func TestTryToConnectTLS(t *testing.T) {
	var (
		dialCount, dialTLSCount int

		err = errors.New("stop connect here")
	)

	dial = func(_ string) (*amqp.Connection, error) {
		dialCount++
		return nil, err
	}

	dialTLS = func(_ string, _ *tls.Config) (*amqp.Connection, error) {
		dialTLSCount++
		return nil, err
	}

	testcases := []struct {
		title     string
		url       string
		secure    bool
		tlsConfig *tls.Config
		wantTLS   bool
	}{
		{"unsecure url, secure false, no tls config", "amqp://example.com", false, nil, false},
		{"secure url, secure false, no tls config", "amqps://example.com", false, nil, true},
		{"unsecure url, secure true, no tls config", "amqp://example.com", true, nil, true},
		{"unsecure url, secure false, tls config", "amqp://example.com", false, &tls.Config{}, true},
	}

	for _, test := range testcases {
		dialCount, dialTLSCount = 0, 0

		conn := newRabbitMQConn("exchange", []string{test.url})
		conn.tryConnect(test.secure, test.tlsConfig)

		have := dialCount
		if test.wantTLS {
			have = dialTLSCount
		}

		if have != 1 {
			t.Errorf("%s: used wrong dialer, Dial called %d times, DialTLS called %d times", test.title, dialCount, dialTLSCount)
		}
	}
}
