package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/hitalos/qr-code-generator/handlers"
	"github.com/hitalos/qr-code-generator/static"
	"github.com/hitalos/qr-code-generator/templates"
)

func main() {
	port := flag.Int("p", 3000, "tcp port to listen")

	flag.Parse()

	mux := createApp()

	listen(mux, *port)
}

func createApp() *chi.Mux {
	engine, err := templates.SetTemplates()
	if err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(6))
	r.HandleFunc("/", handlers.Index(engine))
	r.HandleFunc("/qrcode/svg/{str}", handlers.SVGimage)
	r.HandleFunc("/qrcode/png/{str}", handlers.PNGimage)
	r.Handle("/*", http.FileServer(static.Dir))

	return r
}

func listen(mux *chi.Mux, port int) {
	s := http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Addr:         ":" + strconv.Itoa(port),
		Handler:      mux,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
