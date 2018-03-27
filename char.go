package qrw

import (
	"bytes"
	"io"

	"github.com/rsc/qr"
)

// Use Unicode Block Elements to form the QR Code.
var (
	blockBB = []byte(" ")
	blockWB = []byte("▀")
	blockBW = []byte("▄")
	blockWW = []byte("█")
)

// HalfBlockWriter implements QR Writer use Unicode Block Elements,
// output to text.
type HalfBlockWriter struct {
	Writer
	BlockBB []byte
	BlockWB []byte
	BlockBW []byte
	BlockWW []byte
}

func (w *HalfBlockWriter) writeTopQuietZone(size int) (err error) {
	width := size + w.QuietZone*2
	if w.QuietZone%2 == 0 {
		for i := 0; err == nil && i < width; i++ {
			err = w.Write(w.BlockWW)
		}
	} else {
		for i := 0; err == nil && i < width; i++ {
			err = w.Write(w.BlockBW)
		}
	}
	err = w.Write(lf)
	for j := 3; err == nil && j <= w.QuietZone; j += 2 {
		for i := 0; err == nil && i < width; i++ {
			err = w.Write(w.BlockWW)
		}
		err = w.Write(lf)
	}
	return
}

func (w *HalfBlockWriter) writeBottomQuietZone(size int) (err error) {
	if w.QuietZone <= 1 {
		return nil
	}
	width := size + w.QuietZone*2
	for j := 4; err == nil && j <= w.QuietZone; j += 2 {
		for i := 0; err == nil && i < width; i++ {
			err = w.Write(w.BlockWW)
		}
		err = w.Write(lf)
	}
	if w.QuietZone%2 == 0 {
		for i := 0; err == nil && i < width; i++ {
			err = w.Write(w.BlockWB)
		}
	} else {
		for i := 0; err == nil && i < width; i++ {
			err = w.Write(w.BlockWW)
		}
	}
	err = w.Write(lf)
	return
}

func (w *HalfBlockWriter) writeLeftQuietZone() (err error) {
	for i := 0; err == nil && i < w.QuietZone; i++ {
		err = w.Write(w.BlockWW)
	}
	return
}

func (w *HalfBlockWriter) writeRightQuietZone() (err error) {
	for i := 1; err == nil && i < w.QuietZone; i++ {
		err = w.Write(w.BlockWW)
	}
	err = w.Write(lf)
	return
}

func (w *HalfBlockWriter) writeBlocks(code *qr.Code) error {
	w.init()
	var block uint8
	var bottom bool
	w.writeTopQuietZone(code.Size)
	for y := 0; y <= code.Size; y += 2 {
		w.writeLeftQuietZone()
		bottom = y+1 < code.Size
		for x := 0; x <= code.Size; x++ {
			block = 0
			if code.Black(x, y) {
				block |= 1 << 1
			}
			if bottom && code.Black(x, y+1) {
				block |= 1
			}
			switch block {
			case 3: // 11
				w.Write(w.BlockBB)
			case 2: // 10
				w.Write(w.BlockBW)
			case 1: // 01
				w.Write(w.BlockWB)
			default: // 00
				w.Write(w.BlockWW)
			}
		}
		w.writeRightQuietZone()
	}
	return w.writeBottomQuietZone(code.Size)
}

// QR encode text at the given error correction level,
// and write to the given io.Writer.
func (w *HalfBlockWriter) QR(text string) error {
	code, err := qr.Encode(text, qr.Level(w.Level))
	if err != nil {
		return err
	}
	return w.writeBlocks(code)
}

// QRFile encode text at the given error correction level,
// and rewrite to a file named by filename.
func (w *HalfBlockWriter) QRFile(filename, text string) error {
	buf := &bytes.Buffer{}
	w.Writer.Writer = buf
	if err := w.QR(text); err != nil {
		return err
	}
	return w.WriteFile(filename, buf.Bytes())
}

// NewHalfBlockWriter returns a HalfBlockWriter instance after initialization.
func NewHalfBlockWriter(l Level, w io.Writer) *HalfBlockWriter {
	return &HalfBlockWriter{
		Writer: Writer{
			Level:  l,
			Writer: w,
		},
		BlockBB: blockBB,
		BlockWB: blockWB,
		BlockBW: blockBW,
		BlockWW: blockWW,
	}
}
