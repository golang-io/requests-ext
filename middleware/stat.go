package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/golang-io/requests"
	"net/http"
	"time"
)

func ServeStat(ww *requests.ResponseWriter, r *http.Request, start time.Time, buf *bytes.Buffer) *requests.Stat {
	stat := &requests.Stat{
		StartAt: start.Format("2006-01-02 15:04:05.000"),
		Cost:    time.Since(start).Milliseconds(),
	}
	stat.Request.Remote = r.RemoteAddr
	stat.Request.Method = r.Method
	stat.Request.Header = make(map[string]string)
	for k, v := range r.Header {
		stat.Request.Header[k] = v[0]
	}
	stat.Request.URL = r.URL.String()

	if buf != nil {
		m := make(map[string]any)
		if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
			stat.Request.Body = buf.String()
		} else {
			stat.Request.Body = m
		}
	}
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	stat.Response.URL = scheme + r.Host
	stat.Response.StatusCode = ww.StatusCode
	stat.Response.ContentLength = ww.ContentLength
	stat.Response.Header = make(map[string]string)
	for k, v := range r.Header {
		stat.Response.Header[k] = v[0]
	}
	stat.Response.Body = "[-]"
	return stat
}
