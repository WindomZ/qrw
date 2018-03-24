package qrw

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/rsc/qr"
)

// Level denotes a QR error correction level.
// From least to most tolerant of errors, they are L, M, Q, H.
type Level qr.Level

const (
	// L low tolerant of errors
	L = Level(qr.L) // 20% redundant
	// M middle tolerant of errors
	M = Level(qr.M) // 38% redundant
	// Q half tolerant of errors
	Q = Level(qr.Q) // 55% redundant
	// H high tolerant of errors
	H = Level(qr.H) // 65% redundant
)

// QuietZoneBlocks is the number of QR quiet zone blocks
const QuietZoneBlocks = 4

// CharWrite a QR Code with Unicode Block Elements and write to io.Writer
func CharWrite(w io.Writer, l Level, text string) error {
	return NewHalfBlockWriter(l, w).QR(text)
}

// CharWriteFile a QR Code with Unicode Block Elements and write to file with path.
func CharWriteFile(path string, l Level, text string) error {
	w := &bytes.Buffer{}
	if err := NewHalfBlockWriter(l, w).QR(text); err != nil {
		return err
	}
	return ioutil.WriteFile(path, w.Bytes(), 0666)
}
