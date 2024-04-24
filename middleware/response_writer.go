package middleware

import (
	"fmt"
	"github.com/golang-io/requests"
	"net/http"
	"time"
)

func ServeLog(f func(*requests.Stat)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := &requests.ResponseWriter{ResponseWriter: w}
			buf, body, _ := requests.CopyBody(r.Body)
			r.Body = body
			defer func() {
				f(ServeStat(ww, r, start, buf))
			}()
			next.ServeHTTP(ww, r)
		})
	}
}

// PrintStat is used for server side
func PrintStat(stat *requests.Stat) string {
	return fmt.Sprintf("%s %s \"%s -> %s%s\" - %d %dB in %dms",
		stat.StartAt, stat.Request.Method,
		stat.Request.Remote, stat.Response.URL, stat.Request.URL,
		stat.Response.StatusCode, stat.Response.ContentLength, stat.Cost)
}
