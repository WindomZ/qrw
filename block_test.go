package qrw

import (
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestBlockWriter_QR(t *testing.T) {
	w := NewBlockWriter(L, os.Stdout)
	assert.NoError(t, w.QR("Hello world!"))
}

func TestBlockWriter_Invert(t *testing.T) {
	w := NewBlockWriter(L, os.Stdout)
	assert.NoError(t, w.Invert().QR("Hello world!"))
}

func TestBlockWriter_QRFile(t *testing.T) {
	w := NewBlockWriter(L, os.Stdout)
	assert.NoError(t, w.Invert().QRFile("file_qr_block", "Hello world!"))
}

func BenchmarkBlockWriter_QR(b *testing.B) {
	w := NewBlockWriter(L, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.QR("https://github.com/WindomZ/qrw")
	}
}
