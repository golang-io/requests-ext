package main

import (
	"context"
	"github.com/golang-io/requests"
	"github.com/golang-io/requests-ext/middleware"
	"log"
	"net/http"
)

func main() {
	mux := requests.NewServeMux(requests.URL("0.0.0.0:8081"))
	mux.Use(middleware.ServeLog(func(stat *requests.Stat) {
		log.Printf("%s", middleware.PrintStat(stat))
	}))
	mux.OnStartup(func(s *http.Server) {
		log.Printf("http(s) serve %s", s.Addr)
	})
	mux.OnShutdown(func(s *http.Server) {
		log.Printf("http shutdown.")
	})
	if err := requests.ListenAndServe(context.Background(), mux); err != nil {
		panic(err)
	}
}
