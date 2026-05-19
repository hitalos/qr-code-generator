package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"github.com/hitalos/qr-code-generator/handlers"
	"github.com/hitalos/qr-code-generator/static"
)

const MAX_REQUESTS_PER_MINUTE = 600

func main() {
	port := flag.String("p", ":3000", "tcp port to listen")

	flag.Parse()

	mux := createApp()

	listen(mux, *port)
}

func createApp() *chi.Mux {
	limiter := httprate.LimitByIP(MAX_REQUESTS_PER_MINUTE, time.Minute)
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(limiter)
	r.Use(middleware.Compress(6))
	r.HandleFunc("/", handlers.Index())
	r.HandleFunc("/qrcode/svg/{str}", handlers.SVGimage)
	r.HandleFunc("/qrcode/png/{str}", handlers.PNGimage)
	r.Handle("/*", http.FileServer(static.Dir))

	return r
}

func listen(mux *chi.Mux, port string) {
	s := http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Addr:         port,
		Handler:      mux,
	}

	log.Printf("Listening on: %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
