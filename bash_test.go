package qrw

import (
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestBashWriter_QR(t *testing.T) {
	w := NewBashWriter(L, os.Stdout)
	assert.NoError(t, w.QR("Hello world!"))
}
