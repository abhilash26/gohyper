package middleware

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func FileServer(pattern, folder string) func(http.Handler) http.Handler {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, folder))
	staticHandler := http.StripPrefix(pattern, http.FileServer(filesDir))

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// If the request matches the pattern, serve the static files
			if r.URL.Path == pattern || strings.HasPrefix(r.URL.Path, pattern+"/") {
				staticHandler.ServeHTTP(w, r)
				return
			}
			// Otherwise, continue to the next handler
			next.ServeHTTP(w, r)
		})
	}
}
