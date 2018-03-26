package qrw

import (
	"image"
	"image/color"

	"github.com/rsc/qr"
)

// Image implements image.Image
type Image struct {
	*qr.Code
	halfScale int
}

var (
	whiteColor = color.Gray{0xFF}
	blackColor = color.Gray{0x00}
)

func (c *Image) Bounds() image.Rectangle {
	d := (c.Size + 8) * c.Scale
	return image.Rect(0, 0, d, d)
}

func (c *Image) At(x, y int) color.Color {
	x = x/c.Scale - c.halfScale
	y = y/c.Scale - c.halfScale
	if x < 0 || y < 0 {
		return whiteColor
	}
	if c.Black(x, y) {
		return blackColor
	}
	return whiteColor
}

func (c *Image) ColorModel() color.Model {
	return color.GrayModel
}

func newImage(c *qr.Code) image.Image {
	return &Image{Code: c, halfScale: c.Scale / 2}
}
