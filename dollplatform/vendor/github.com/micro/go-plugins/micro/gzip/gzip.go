// Package gzip is a micro plugin for gzipping http response
package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"
)

type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

type gzipper struct{}

func (g *gzipper) Flags() []cli.Flag {
	return nil
}

func (g *gzipper) Commands() []cli.Command {
	return nil
}

func (g *gzipper) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// has gzip accept-encoding?
			if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				h.ServeHTTP(w, r)
				return
			}

			// set the content-encoding
			w.Header().Set("Content-Encoding", "gzip")

			// create gzip writer
			gz := gzip.NewWriter(w)
			defer gz.Close()

			// create http response writer
			gzw := gzipWriter{gz, w}

			// serve the request
			h.ServeHTTP(gzw, r)
		})
	}
}

func (g *gzipper) Init(ctx *cli.Context) error {
	return nil
}

func (g *gzipper) String() string {
	return "gzip"
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func New() plugin.Plugin {
	return new(gzipper)
}
