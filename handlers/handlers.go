package handlers

import (
	"image/png"
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/template/html"
)

func Index(engine *html.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := engine.Render(w, "index", nil, "layouts/main"); err != nil {
			log.Println(err)
		}
	}
}

func PNGimage(w http.ResponseWriter, r *http.Request) {
	str := chi.URLParam(r, "str")
	qr, _ := qr.Encode(str, qr.M, qr.Auto)
	size := qr.Bounds().Max.X * 20
	qr, _ = barcode.Scale(qr, size, size)

	if err := png.Encode(w, qr); err != nil {
		log.Println(err)
	}
}

func SVGimage(w http.ResponseWriter, r *http.Request) {
	str := chi.URLParam(r, "str")
	qr, err := qr.Encode(str, qr.M, qr.Auto)
	if err != nil {
		log.Println(err)
	}

	qs := qrSVG{
		qr:        qr,
		blockSize: 10,
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	defer s.End()

	if err = qs.WriteQrSVG(s); err != nil {
		log.Println(err)
	}
}
