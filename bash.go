package qrw

import (
	"bytes"
	"io"
	"os/exec"

	"github.com/rsc/qr"
)

// Use two Unicode characters to form the QR Code.
var (
	blockBashB = []byte("\033[47;30m  \033[0m")
	blockBashW = []byte("\033[40;37m  \033[0m")
)

// BashWriter implements QR Writer by 'echo' command,
// output to terminal stdout.
type BashWriter struct {
	Writer
	Buffer *bytes.Buffer
	BlockB []byte
	BlockW []byte
}

// Write the io.Writer wraps the basic Write method.
func (w *BashWriter) Write(p []byte) (err error) {
	if w.Buffer == nil {
		return w.Writer.Write(p)
	}
	_, err = w.Buffer.Write(p)
	return
}

func (w *BashWriter) echo() (err error) {
	cmd := exec.Command("echo", "-e", w.Buffer.String())
	cmd.Stdout = w.Writer.Writer
	cmd.Stderr = w.Writer.Writer
	err = cmd.Run()
	return
}

func (w *BashWriter) writeTopQuietZone(size int) (err error) {
	width := size + QuietZoneBlocks*2
	for i := 0; err == nil && i < 4; i++ {
		for i := 0; i < width; i++ {
			w.Write(w.BlockW)
		}
		err = w.Write(lf)
	}
	return
}

func (w *BashWriter) writeBottomQuietZone(size int) (err error) {
	width := size + QuietZoneBlocks*2
	for i := 1; err == nil && i < 4; i++ {
		for i := 0; i < width; i++ {
			w.Write(w.BlockW)
		}
		err = w.Write(lf)
	}
	return
}

func (w *BashWriter) writeLeftQuietZone() (err error) {
	for i := 0; i < QuietZoneBlocks; i++ {
		err = w.Write(w.BlockW)
	}
	return
}

func (w *BashWriter) writeRightQuietZone() error {
	for i := 1; i < QuietZoneBlocks; i++ {
		w.Write(w.BlockW)
	}
	return w.Write(lf)
}

func (w *BashWriter) writeBlocks(code *qr.Code) error {
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
	w.writeBottomQuietZone(code.Size)
	return w.echo()
}

// QR encode text at the given error correction level,
// and write to the given io.Writer.
func (w *BashWriter) QR(text string) error {
	code, err := qr.Encode(text, qr.Level(w.Level))
	if err != nil {
		return err
	}
	return w.writeBlocks(code)
}

// NewBashWriter returns a BashWriter instance after initialization.
func NewBashWriter(l Level, w io.Writer) *BashWriter {
	return &BashWriter{
		Writer: Writer{
			Level:  l,
			Writer: w,
		},
		Buffer: &bytes.Buffer{},
		BlockB: blockBashB,
		BlockW: blockBashW,
	}
}
