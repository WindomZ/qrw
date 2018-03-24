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

func BenchmarkHalfBlockWriter_QR(b *testing.B) {
	w := NewHalfBlockWriter(L, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.QR("https://github.com/WindomZ/qrw")
	}
}
