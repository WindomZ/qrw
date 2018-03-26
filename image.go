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

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (c *Image) Bounds() image.Rectangle {
	d := (c.Size + 8) * c.Scale
	return image.Rect(0, 0, d, d)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
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

// ColorModel returns the Image's color model.
func (c *Image) ColorModel() color.Model {
	return color.GrayModel
}

func newImage(c *qr.Code) image.Image {
	return &Image{Code: c, halfScale: c.Scale / 2}
}
