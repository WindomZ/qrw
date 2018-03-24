package qrw

import (
	"io"

	"github.com/rsc/qr"
)

// Use two Unicode characters to form the QR Code.
var (
	blockB = []byte("  ")
	blockW = []byte("██")
)

// BlockWriter implements QR Writer use Unicode characters,
// output to text.
type BlockWriter struct {
	Writer
	BlockB []byte
	BlockW []byte
}

func (w *BlockWriter) writeTopQuietZone(size int) (err error) {
	width := size + QuietZoneBlocks*2
	for i := 0; err == nil && i < 4; i++ {
		for i := 0; i < width; i++ {
			w.Write(w.BlockW)
		}
		err = w.Write(lf)
	}
	return
}

func (w *BlockWriter) writeBottomQuietZone(size int) (err error) {
	width := size + QuietZoneBlocks*2
	for i := 1; err == nil && i < 4; i++ {
		for i := 0; i < width; i++ {
			w.Write(w.BlockW)
		}
		err = w.Write(lf)
	}
	return
}

func (w *BlockWriter) writeLeftQuietZone() (err error) {
	for i := 0; i < QuietZoneBlocks; i++ {
		err = w.Write(w.BlockW)
	}
	return
}

func (w *BlockWriter) writeRightQuietZone() error {
	for i := 1; i < QuietZoneBlocks; i++ {
		w.Write(w.BlockW)
	}
	return w.Write(lf)
}

func (w *BlockWriter) writeBlocks(code *qr.Code) error {
	w.writeTopQuietZone(code.Size)
	for i := 0; i <= code.Size; i++ {
		w.writeLeftQuietZone()
		for j := 0; j <= code.Size; j++ {
			if code.Black(i, j) {
				w.Write(w.BlockB)
			} else {
				w.Write(w.BlockW)
			}
		}
		w.writeRightQuietZone()
	}
	return w.writeBottomQuietZone(code.Size)
}

// QR encode text at the given error correction level,
// and write to the given io.Writer.
func (w *BlockWriter) QR(text string) error {
	code, err := qr.Encode(text, qr.Level(w.Level))
	if err != nil {
		return err
	}
	return w.writeBlocks(code)
}

// NewBlockWriter returns a BlockWriter instance after initialization.
func NewBlockWriter(l Level, w io.Writer) *BlockWriter {
	return &BlockWriter{
		Writer: Writer{
			Level:  l,
			Writer: w,
		},
		BlockB: blockB,
		BlockW: blockW,
	}
}
