package qrw

import (
	"bytes"
	"image/jpeg"

	"github.com/rsc/qr"
)

// JPEGWriter implements QR Writer for a JPEG image.
type JPEGWriter struct {
	Writer
}

// QRFile encode text at the given error correction level,
// and rewrite to a file named by filename.
func (w *JPEGWriter) QRFile(filename, text string) error {
	code, err := qr.Encode(text, qr.Level(w.Level))
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	if err = jpeg.Encode(buf, newImage(code),
		&jpeg.Options{Quality: 100}); err != nil {
		return err
	}
	return w.WriteFile(filename, buf.Bytes())
	//return nil
}

// NewJPEGWriter returns a JPEGWriter instance after initialization.
func NewJPEGWriter(l Level) *JPEGWriter {
	return &JPEGWriter{
		Writer: Writer{
			Level: l,
		},
	}
}
