package qrw

import "github.com/rsc/qr"

// PNGWriter implements QR Writer for a PNG image.
type PNGWriter struct {
	Writer
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
func NewPNGWriter(l Level) *PNGWriter {
	return &PNGWriter{
		Writer: Writer{
			Level: l,
		},
	}
}
