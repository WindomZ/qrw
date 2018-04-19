package qrw

import (
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestJPEGWriter_QR(t *testing.T) {
	f, err := os.OpenFile("file_qr_1.jpeg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	w := NewJPEGWriter(L, f)
	assert.NoError(t, w.QR("Hello world!"))
}

func TestJPEGWriter_QRFile(t *testing.T) {
	w := NewJPEGWriter(L)
	assert.NoError(t, w.QRFile("file_qr_2.jpeg", "Hello world!"))

	assert.Error(t, JPEG("test/file_qr.jpeg", L, "Hello world!"))
}
