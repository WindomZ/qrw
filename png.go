package qrw

import (
	"io"

	"github.com/rsc/qr"
)

// PNGWriter implements QR Writer for a PNG image.
type PNGWriter struct {
	Writer
}

// QR encode text at the given error correction level,
// and write to the given io.Writer.
func (w *PNGWriter) QR(text string) error {
	code, err := qr.Encode(text, qr.Level(w.Level))
	if err != nil {
		return err
	}
	return w.Write(code.PNG())
}

// QRFile encode text at the given error correction level,
// and rewrite to a file named by filename.
func (w *PNGWriter) QRFile(filename, text string) error {
	code, err := qr.Encode(text, qr.Level(w.Level))
	if err != nil {
		return err
	}
	return w.WriteFile(filename, code.PNG())
}

// NewPNGWriter returns a PNGWriter instance after initialization.
func NewPNGWriter(l Level, ws ...io.Writer) *PNGWriter {
	var w io.Writer
	if len(ws) != 0 {
		w = ws[0]
	}
	return &PNGWriter{
		Writer: Writer{
			Level:  l,
			Writer: w,
		},
	}
}
