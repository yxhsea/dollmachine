package http

import (
	"net/http"
)

type AuthTransport struct {
	http.RoundTripper
	auth string
}

func NewAuthTransport(basicAuth string, transport http.RoundTripper) http.RoundTripper {
	return &AuthTransport{transport, basicAuth}
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.auth)
	return t.RoundTripper.RoundTrip(req)
}
