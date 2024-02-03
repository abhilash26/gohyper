package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

func GzipCompression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if client supports gzip encoding
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// Set Content-Encoding header to inform the client about gzip encoding
			w.Header().Set("Content-Encoding", "gzip")

			// Create a gzip writer and wrap the ResponseWriter
			gz := gzip.NewWriter(w)
			defer gz.Close()
			w = &gzipResponseWriter{Writer: gz, ResponseWriter: w}
		}

		// Continue to the next handler
		next.ServeHTTP(w, r)
	})
}

// Custom ResponseWriter to handle gzip compression
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *gzipResponseWriter) Flush() {
	w.Writer.Flush()
}

func (w *gzipResponseWriter) WriteHeader(code int) {
	w.Header().Del("Content-Length") // Remove Content-Length header when using gzip
	w.ResponseWriter.WriteHeader(code)
}
