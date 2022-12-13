package handlers

import (
	"errors"
	"fmt"
	"image/color"

	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode"
)

type qrSVG struct {
	qr        barcode.Barcode
	qrWidth   int
	blockSize int
	startingX int
	startingY int
}

func (qs *qrSVG) WriteQrSVG(s *svg.SVG) error {
	qs.startingX = 2
	qs.startingY = 2
	qs.qrWidth = qs.qr.Bounds().Max.X
	width := qs.qrWidth + 4

	s.Startraw(fmt.Sprintf("viewBox=\"0 0 %d %[1]d\"", width*qs.blockSize))
	s.Style("text/css", "rect { fill: white }")

	if qs.qr.Metadata().CodeKind == "QR Code" {
		currY := qs.startingY

		s.Group(fmt.Sprintf("transform: scale(%d)", qs.blockSize))
		for x := 0; x < qs.qrWidth; x++ {
			currX := qs.startingX
			for y := 0; y < qs.qrWidth; y++ {
				if qs.qr.At(x, y) == color.Black {
					s.Rect(currX, currY, 1, 1, "fill:black")
				} else if qs.qr.At(x, y) == color.White {
					s.Rect(currX, currY, 1, 1)
				}
				currX += 1
			}
			currY += 1
		}
		s.Gend()
		return nil
	}
	return errors.New("can not write to SVG: Not a QR code")
}
