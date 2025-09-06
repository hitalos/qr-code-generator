package handlers

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"net/url"

	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/template/html/v2"
)

func Index(engine *html.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := engine.Render(w, "index", nil, "layouts/main"); err != nil {
			errHandler(err, w)
		}
	}
}

func createBarcode(r *http.Request, scale int) (barcode.Barcode, error) {
	str, err := url.QueryUnescape(chi.URLParam(r, "str"))
	if err != nil {
		return nil, err
	}

	qrcode, err := qr.Encode(str, qr.L, qr.Auto)
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

	border := 20
	bColor := color.RGBA{255, 255, 255, 255}
	borderedImage := image.NewRGBA(image.Rect(0, 0, qrcode.Bounds().Dx()+2*border, qrcode.Bounds().Dy()+2*border))
	draw.Draw(borderedImage, borderedImage.Bounds(), &image.Uniform{C: bColor}, image.Point{}, draw.Src)
	draw.Draw(borderedImage, image.Rect(border, border, qrcode.Bounds().Dx()+border, qrcode.Bounds().Dy()+border), qrcode, image.Point{}, draw.Src)

	if err = png.Encode(w, borderedImage); err != nil {
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
	_, _ = w.Write([]byte("error!"))
}
