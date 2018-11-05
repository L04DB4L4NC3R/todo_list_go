package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipCompression struct {
	Next http.Handler
}

func (g GzipCompression) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if g.Next == nil {
		g.Next = http.DefaultServeMux
	}
	encodings := r.Header.Get("Accept-Encoding")

	if !strings.Contains(encodings, "gzip") {
		g.ServeHTTP(w, r)
		return
	}

	w.Header().Add("Content-Encoding", "gzip")
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	grw := gzipResponseWriter{
		ResponseWriter: w,
		Writer:         gzipWriter}
	g.ServeHTTP(grw, r)

}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (grw gzipResponseWriter) Write(b []byte) (int, error) {
	return grw.Writer.Write(b)
}
