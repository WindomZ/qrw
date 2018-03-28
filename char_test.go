package qrw

import (
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestHalfBlockWriter_QR(t *testing.T) {
	w := NewHalfBlockWriter(L, os.Stdout)
	assert.NoError(t, w.QR("Hello world!"))
}

func TestHalfBlockWriter_Invert(t *testing.T) {
	w := NewHalfBlockWriter(L, os.Stdout)
	assert.NoError(t, w.Invert().QR("Hello world!"))
}

func TestHalfBlockWriter_QRFile(t *testing.T) {
	w := NewHalfBlockWriter(L, os.Stdout)
	assert.NoError(t, w.Invert().QRFile("file_qr_half_block", "Hello world!"))
}

func BenchmarkHalfBlockWriter_QR(b *testing.B) {
	w := NewHalfBlockWriter(L, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.QR("https://github.com/WindomZ/qrw")
	}
}
