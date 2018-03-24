package qrw

import (
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

func (w *HalfBlockWriter) writeTopQuietZone(size int) error {
	width := size + QuietZoneBlocks*2
	for i := 0; i < width; i++ {
		w.Write(w.BlockWW)
	}
	w.Write(lf)
	for i := 0; i < width; i++ {
		w.Write(w.BlockWW)
	}
	return w.Write(lf)
}

func (w *HalfBlockWriter) writeBottomQuietZone(size int) error {
	width := size + QuietZoneBlocks*2
	for i := 0; i < width; i++ {
		w.Write(w.BlockWW)
	}
	w.Write(lf)
	for i := 0; i < width; i++ {
		w.Write(w.BlockWB)
	}
	return w.Write(lf)
}

func (w *HalfBlockWriter) writeLeftQuietZone() (err error) {
	for i := 0; i < QuietZoneBlocks; i++ {
		err = w.Write(w.BlockWW)
	}
	return
}

func (w *HalfBlockWriter) writeRightQuietZone() error {
	for i := 1; i < QuietZoneBlocks; i++ {
		w.Write(w.BlockWW)
	}
	return w.Write(lf)
}

func (w *HalfBlockWriter) writeBlocks(code *qr.Code) error {
	var block uint8
	var bottom bool
	w.writeTopQuietZone(code.Size)
	for i := 0; i <= code.Size; i += 2 {
		w.writeLeftQuietZone()
		bottom = i+1 < code.Size
		for j := 0; j <= code.Size; j++ {
			block = 0
			if code.Black(i, j) {
				block |= 1 << 1
			}
			if bottom && code.Black(i+1, j) {
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
