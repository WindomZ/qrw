package qrw

import (
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestHalfBashWriter_QR(t *testing.T) {
	w := NewHalfBashWriter(L, os.Stdout)
	assert.NoError(t, w.QR("Hello world!"))
}

func TestHalfBashWriter_Invert(t *testing.T) {
	w := NewHalfBashWriter(L, os.Stdout)
	assert.NoError(t, w.Invert().QR("Hello world!"))
}
