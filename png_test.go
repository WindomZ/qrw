package qrw

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestPNGWriter_QR(t *testing.T) {
	w := NewPNGWriter(L)
	assert.NoError(t, w.QRFile("file_qr.png", "Hello world!"))
	assert.Error(t, CharWriteFile("test/file_qr.png", L, "Hello world!"))
}
