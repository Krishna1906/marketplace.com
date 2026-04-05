package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		ip := r.RemoteAddr

		// ---------- ANSI COLORS ----------
		reset := "\033[0m"

		// Background colors
		bgBlue := "\033[44m"   // GET
		bgGreen := "\033[42m"  // POST
		bgYellow := "\033[43m" // PUT
		bgRed := "\033[41m"    // DELETE

		// Foreground
		fgWhite := "\033[97m"

		methodBg := bgBlue
		switch r.Method {
		case http.MethodPost:
			methodBg = bgGreen
		case http.MethodPut:
			methodBg = bgYellow
		case http.MethodDelete:
			methodBg = bgRed
		}

		// Status background
		statusBg := bgGreen
		if rw.status >= 400 {
			statusBg = bgYellow
		}
		if rw.status >= 500 {
			statusBg = bgRed
		}

		log.Printf(
			"[HTTP] %s | %s%s %d %s | %v | %s | %s%s %-6s %s \"%s\"",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusBg, fgWhite, rw.status, reset,
			duration,
			ip,
			methodBg, fgWhite, r.Method, reset,
			r.URL.Path,
		)
	})
}
