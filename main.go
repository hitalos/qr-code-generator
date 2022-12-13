package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
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

	log.Printf("Listening on: \n\thttp://%s%s", strings.Join(localAddresses(), s.Addr+"\n\thttp://"), s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func localAddresses() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
	}

	ips := []string{}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatalln(err)
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				if v.IP.To4() == nil {
					ips = append(ips, fmt.Sprintf("[%s]", v.IP))
					continue
				}
				ips = append(ips, v.IP.String())
			}
		}
	}

	return ips
}
