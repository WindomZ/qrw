package qrw

import (
	"bytes"
	"io"
	"os/exec"

	"github.com/rsc/qr"
)

// Use two Unicode characters to form the QR Code.
var (
	blockBashBB = []byte("\u001b[30m\u001b[40m\u2588\u001b[0m")
	blockBashBW = []byte("\u001b[37m\u001b[40m\u2584\u001b[0m")
	blockBashWB = []byte("\u001b[30m\u001b[47m\u2585\u001b[0m")
	blockBashWW = []byte("\u001b[37m\u001b[47m\u2588\u001b[0m")
)

// HalfBashWriter implements QR Writer by 'echo' command,
// output to terminal stdout.
type HalfBashWriter struct {
	Writer
	Buffer  *bytes.Buffer
	BlockBB []byte
	BlockWB []byte
	BlockBW []byte
	BlockWW []byte
}

// Write the io.Writer wraps the basic Write method.
func (w *HalfBashWriter) Write(p []byte) (err error) {
	if w.Buffer == nil {
		return w.Writer.Write(p)
	}
	_, err = w.Buffer.Write(p)
	return
}

func (w *HalfBashWriter) echo() (err error) {
	cmd := exec.Command("echo", "-e", w.Buffer.String())
	cmd.Stdout = w.Writer.Writer
	cmd.Stderr = w.Writer.Writer
	err = cmd.Run()
	return
}

func (w *HalfBashWriter) writeTopQuietZone(size int) (err error) {
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

func (w *HalfBashWriter) writeBottomQuietZone(size int) (err error) {
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

func (w *HalfBashWriter) writeLeftQuietZone() (err error) {
	for i := 0; err == nil && i < w.QuietZone; i++ {
		err = w.Write(w.BlockWW)
	}
	return
}

func (w *HalfBashWriter) writeRightQuietZone() (err error) {
	for i := 1; err == nil && i < w.QuietZone; i++ {
		err = w.Write(w.BlockWW)
	}
	err = w.Write(lf)
	return
}
func (w *HalfBashWriter) writeBlocks(code *qr.Code) error {
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
	w.writeBottomQuietZone(code.Size)
	return w.echo()
}

// Invert invert the color of block.
func (w *HalfBashWriter) Invert() *HalfBashWriter {
	w.BlockBB, w.BlockWB, w.BlockBW, w.BlockWW = w.BlockWW, w.BlockBW, w.BlockWB, w.BlockBB
	return w
}

// QR encode text at the given error correction level,
// and write to the given io.Writer.
func (w *HalfBashWriter) QR(text string) error {
	code, err := qr.Encode(text, qr.Level(w.Level))
	if err != nil {
		return err
	}
	return w.writeBlocks(code)
}

// NewHalfBashWriter returns a HalfBashWriter instance after initialization.
func NewHalfBashWriter(l Level, w io.Writer) *HalfBashWriter {
	return &HalfBashWriter{
		Writer: Writer{
			Level:  l,
			Writer: w,
		},
		Buffer:  &bytes.Buffer{},
		BlockBB: blockBashBB,
		BlockWB: blockBashWB,
		BlockBW: blockBashBW,
		BlockWW: blockBashWW,
	}
}
