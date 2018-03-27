package qrw

import (
	"io"
	"io/ioutil"
)

// line break
var lf = []byte("\n")

// Writer denotes a basic QR io.Writer.
type Writer struct {
	Level     Level
	Writer    io.Writer
	QuietZone int
}

// initial Writer
func (w *Writer) init() *Writer {
	if w.QuietZone <= 0 {
		w.QuietZone = QuietZoneBlocks
	}
	return w
}

// Write the io.Writer wraps the basic Write method.
func (w *Writer) Write(p []byte) (err error) {
	if w.Writer == nil {
		return
	}
	_, err = w.Writer.Write(p)
	return
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm;
// otherwise WriteFile truncates it before writing.
func (w *Writer) WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0666)
}
