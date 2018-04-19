package qrw

import (
	"bytes"
	"image/jpeg"
	"io"

	"github.com/rsc/qr"
)

// JPEGWriter implements QR Writer for a JPEG image.
type JPEGWriter struct {
	Writer
}

// QR encode text at the given error correction level,
// and write to the given io.Writer.
func (w *JPEGWriter) QR(text string) error {
	code, err := qr.Encode(text, qr.Level(w.Level))
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	if err = jpeg.Encode(buf, newImage(code),
		&jpeg.Options{Quality: 100}); err != nil {
		return err
	}
	return w.Write(buf.Bytes())
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
func NewJPEGWriter(l Level, ws ...io.Writer) *JPEGWriter {
	var w io.Writer
	if len(ws) != 0 {
		w = ws[0]
	}
	return &JPEGWriter{
		Writer: Writer{
			Level:  l,
			Writer: w,
		},
	}
}
