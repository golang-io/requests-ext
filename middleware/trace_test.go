package middleware_test

import (
	"context"
	"github.com/golang-io/requests"
	"github.com/golang-io/requests-ext/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_Trace(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, _ = w.Write([]byte("hello world"))
	}))
	defer s.Close()

	sess := requests.New()
	resp, err := sess.DoRequest(
		context.Background(),
		requests.Timeout(600*time.Minute),
		requests.URL(s.URL),
		requests.Logf(func(ctx context.Context, stat *requests.Stat) {
			t.Logf("%s", stat)
		}),
		middleware.TraceLv(5, 10000),
	)
	if err != nil {
		t.Logf("%v", err)
		panic(err)
		return
	}
	t.Logf("%#v", resp.Content.String())
}
