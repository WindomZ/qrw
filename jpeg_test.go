package qrw

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestJPEGWriter_QR(t *testing.T) {
	w := NewJPEGWriter(L)
	assert.NoError(t, w.QRFile("file_qr.jpeg", "Hello world!"))
	assert.Error(t, CharWriteFile("test/file_qr.jpeg", L, "Hello world!"))
}
