package qrw

import (
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestPNGWriter_QR(t *testing.T) {
	f, err := os.OpenFile("file_qr_1.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	w := NewPNGWriter(L, f)
	assert.NoError(t, w.QR("Hello world!"))
}

func TestPNGWriter_QRFile(t *testing.T) {
	w := NewPNGWriter(L)
	assert.NoError(t, w.QRFile("file_qr_2.png", "Hello world!"))

	assert.Error(t, PNG("test/file_qr.png", L, "Hello world!"))
}
