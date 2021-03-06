package qrw

import (
	"io"
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

// QuietZoneBlocks is the default number of QR quiet zone blocks
const QuietZoneBlocks = 4

// HalfBlockWrite generates a QR Code with Unicode Block Elements and output to io.Writer.
// Half character as a QR block.
func HalfBlockWrite(w io.Writer, l Level, text string) error {
	return NewHalfBlockWriter(l, w).QR(text)
}

// HalfBlockWriteFile generates a QR Code with Unicode Block Elements and output to a file named by filename.
// Half character as a QR block.
func HalfBlockWriteFile(filename string, l Level, text string) error {
	return NewHalfBlockWriter(l, nil).QRFile(filename, text)
}

// BlockWrite generates a QR Code with Unicode characters and output to io.Writer.
// Two characters as a QR block.
func BlockWrite(w io.Writer, l Level, text string) error {
	return NewBlockWriter(l, w).QR(text)
}

// BlockWriteFile generates a QR Code with Unicode characters and output to a file named by filename.
// Two characters as a QR block.
func BlockWriteFile(filename string, l Level, text string) error {
	return NewBlockWriter(l, nil).QRFile(filename, text)
}

// Bash generates a QR Code with bash color and output to Unix shell.
// Two characters as a QR block.
func Bash(l Level, text string) error {
	return NewBashWriter(l, os.Stdout).QR(text)
}

// HalfBash generates a QR Code with bash color and output to Unix shell.
// Half character as a QR block.
func HalfBash(l Level, text string) error {
	return NewHalfBashWriter(l, os.Stdout).QR(text)
}

// PNG generates a QR Code and output to a PNG image file named by filename.
func PNG(filename string, l Level, text string) error {
	return NewPNGWriter(l).QRFile(filename, text)
}

// JPEG generates a QR Code and output to a JPEG image file named by filename.
func JPEG(filename string, l Level, text string) error {
	return NewJPEGWriter(l).QRFile(filename, text)
}
