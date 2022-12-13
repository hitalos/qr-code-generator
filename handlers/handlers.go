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
			errHandler(err, w)
		}
	}
}

func createBarcode(r *http.Request, scale int) (barcode.Barcode, error) {
	qrcode, err := qr.Encode(chi.URLParam(r, "str"), qr.M, qr.Auto)
	if err != nil {
		return qrcode, err
	}

	if scale > 1 {
		size := qrcode.Bounds().Max.X * scale
		qrcode, err = barcode.Scale(qrcode, size, size)
	}

	return qrcode, err
}

func PNGimage(w http.ResponseWriter, r *http.Request) {
	qrcode, err := createBarcode(r, 20)
	if err != nil {
		errHandler(err, w)
	}

	if err = png.Encode(w, qrcode); err != nil {
		errHandler(err, w)
	}
}

func SVGimage(w http.ResponseWriter, r *http.Request) {
	qrcode, err := createBarcode(r, 1)
	if err != nil {
		errHandler(err, w)
	}

	qs := qrSVG{
		qr:        qrcode,
		blockSize: 10,
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	defer s.End()

	if err := qs.WriteQrSVG(s); err != nil {
		errHandler(err, w)
	}
}

func errHandler(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("error!"))
}
