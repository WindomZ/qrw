package qrw

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

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
var QuietZoneBlocks = 4

// CharWrite generates a QR Code with Unicode Block Elements and output to io.Writer.
// Half character as a QR block.
func CharWrite(w io.Writer, l Level, text string) error {
	return NewHalfBlockWriter(l, w).QR(text)
}

// CharWriteFile generates a QR Code with Unicode Block Elements and output to file with path.
// Half character as a QR block.
func CharWriteFile(path string, l Level, text string) error {
	w := &bytes.Buffer{}
	if err := NewHalfBlockWriter(l, w).QR(text); err != nil {
		return err
	}
	return ioutil.WriteFile(path, w.Bytes(), 0666)
}

// BlockWrite generates a QR Code with Unicode characters and output to io.Writer.
// Two characters as a QR block.
func BlockWrite(w io.Writer, l Level, text string) error {
	return NewBlockWriter(l, w).QR(text)
}

// BlockWriteFile generates a QR Code with Unicode characters and output to file with path.
// Two characters as a QR block.
func BlockWriteFile(path string, l Level, text string) error {
	w := &bytes.Buffer{}
	if err := NewBlockWriter(l, w).QR(text); err != nil {
		return err
	}
	return ioutil.WriteFile(path, w.Bytes(), 0666)
}

// Bash generates a QR Code with bash color and output to Unix shell.
// Two characters as a QR block.
func Bash(l Level, text string) error {
	return NewBashWriter(l, os.Stdout).QR(text)
}
