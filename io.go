package qrw

import "io"

// line break
var lf = []byte("\n")

// Writer denotes a basic QR io.Writer.
type Writer struct {
	Level  Level
	Writer io.Writer
}

// Write the io.Writer wraps the basic Write method.
func (w *Writer) Write(p []byte) (err error) {
	if w.Writer == nil {
		return
	}
	_, err = w.Writer.Write(p)
	return
}
